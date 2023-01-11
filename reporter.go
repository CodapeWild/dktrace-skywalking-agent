package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"sync"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
)

type GRPCReporterRelay struct {
	reporters []go2sky.Reporter
	trace     []go2sky.ReportedSpan
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

	return &GRPCReporterRelay{reporters: reporters}, nil
}

func (r *GRPCReporterRelay) Boot(service string, serviceInstance string, cdsWatchers []go2sky.AgentConfigChangeWatcher) {
	for i := range r.reporters {
		r.reporters[i].Boot(fmt.Sprintf("%s-%d", service, i), fmt.Sprintf("%s:%d", serviceInstance, i), cdsWatchers)
	}
}

func (r *GRPCReporterRelay) Send(spans []go2sky.ReportedSpan) {
	r.trace = append(r.trace, spans...)
	if len(r.trace) == spanCount {
		log.Println("all spans received, ready to send")

		r.sendHelper()
		// close(globalCloser)
	}
}

func (r *GRPCReporterRelay) sendHelper() {
	wg := &sync.WaitGroup{}
	wg.Add(cfg.Sender.Threads)
	for i := range r.reporters {
		dupli := make([]go2sky.ReportedSpan, len(r.trace))
		for j := range r.trace {
			dupli[j] = fromReportedSpan(r.trace[j])
		}

		go func(index int, trace []go2sky.ReportedSpan) {
			defer wg.Done()

			modifyTraceID(trace)

			for j := 0; j < cfg.Sender.SendCount; j++ {
				r.reporters[index].Send(trace)
				log.Printf("reporter[%d] finished %dth send\n", index, j+1)
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
	// oldnew := make(map[string]string)
	buf := make([]byte, 15)
	rand.Read(buf)
	newtid := hex.EncodeToString(buf)
	// oldtid := trace[0].Context().TraceID
	// oldnew[oldtid] = newtid
	for i := range trace {
		trace[i].(*defSpan).ctx.TraceID = newtid
	}
	// for i := range trace {
	// 	for j := range trace[i].Refs() {
	// 		if newtid, ok := oldnew[trace[i].Refs()[j].TraceID]; ok {
	// 			trace[i].Refs()[j].TraceID = newtid
	// 		}
	// 	}
	// }
}
