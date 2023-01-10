package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
)

type GRPCReporterRelay struct {
	reporters []go2sky.Reporter
}

func NewGRPCReporterRelay() (go2sky.Reporter, error) {
	var (
		reporters = make([]go2sky.Reporter, cfg.Sender.Threads)
		err       error
	)
	for i := range reporters {
		if reporters[i], err = reporter.NewGRPCReporter(cfg.DkAgent); err != nil {
			log.Printf("### start reporter[%d] failed", i)

			return nil, err
		}
	}

	return &GRPCReporterRelay{reporters}, nil
}

func (r *GRPCReporterRelay) Boot(service string, serviceInstance string, cdsWatchers []go2sky.AgentConfigChangeWatcher) {
	for i := range r.reporters {
		r.reporters[i].Boot(fmt.Sprintf("%s-%d", service, i), fmt.Sprintf("%s:%d", serviceInstance, i), cdsWatchers)
	}
}

func (r *GRPCReporterRelay) Send(spans []go2sky.ReportedSpan) {
	wg := &sync.WaitGroup{}
	wg.Add(cfg.Sender.Threads)
	for i := range r.reporters {
		dupli := make([]go2sky.ReportedSpan, len(spans))
		for k := range spans {
			dupli[k] = fromReportedSpan(spans[k])
		}

		go func(index int, spans []go2sky.ReportedSpan) {
			defer wg.Done()

			for k := range spans {
				spans[k].(*defSpan).ReNewTraceID()
			}

			for j := 0; j < cfg.Sender.SendCount; j++ {
				r.reporters[index].Send(spans)
				log.Printf("reporter[%d] finished send %d\n", index, j)
			}
		}(i, dupli)
	}
	wg.Wait()
}

func (r *GRPCReporterRelay) Close() {
	for i := range r.reporters {
		r.reporters[i].Close()
	}
}

func modifyTraceID(trace []go2sky.ReportedSpan) {
	oldnew := make(map[string]string)
	for i := range trace {
		old := trace[i].Context().TraceID
		oldnew[old] = trace[i].(*defSpan).ReNewTraceID()
	}
	for i := range trace {
		for j := range trace[i].Refs() {
			if newtid, ok := oldnew[trace[i].Refs()[j].TraceID]; ok {
				trace[i].Refs()[j].TraceID = newtid
			}
		}
	}
}
