package main

import (
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
		r.reporters[i].Boot(service, serviceInstance, cdsWatchers)
	}
}

func (r *GRPCReporterRelay) Send(spans []go2sky.ReportedSpan) {
	wg := &sync.WaitGroup{}
	wg.Add(cfg.Sender.Threads)
	for i := range r.reporters {
		go func(r go2sky.Reporter, spans []go2sky.ReportedSpan) {
			defer wg.Done()

			r.Send(spans)
		}(r.reporters[i], spans)
	}
	wg.Wait()

	close(globalCloser)
}

func (r *GRPCReporterRelay) Close() {
	for i := range r.reporters {
		r.reporters[i].Close()
	}
}
