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
	// for i := range r.reporters {
	// 	log.Println(service, serviceInstance)
	// 	r.reporters[i].Boot(service, serviceInstance, cdsWatchers)
	// }
	for i := range r.reporters {
		r.reporters[i].Boot(fmt.Sprintf("%s-%d", service, i), fmt.Sprintf("%s:%d", serviceInstance, i), cdsWatchers)
	}
}

func (r *GRPCReporterRelay) Send(spans []go2sky.ReportedSpan) {
	wg := &sync.WaitGroup{}
	wg.Add(cfg.Sender.Threads)
	for i := range r.reporters {
		go func(index int, spans []go2sky.ReportedSpan) {
			defer wg.Done()

			for j := 0; j < cfg.Sender.SendCount; j++ {
				r.reporters[index].Send(spans)
				log.Printf("reporter[%d] finished send %d\n", index, j)
			}
		}(i, spans)
	}
	wg.Wait()
}

func (r *GRPCReporterRelay) Close() {
	for i := range r.reporters {
		r.reporters[i].Close()
	}
}
