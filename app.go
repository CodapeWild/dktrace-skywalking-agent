package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/SkyAPM/go2sky"
)

var (
	cfg          *config
	globalCloser = make(chan struct{})
	agentAddress = "127.0.0.1:"
)

type sender struct {
	Threads      int `json:"threads"`
	SendCount    int `json:"send_count"`
	SendInterval int `json:"send_interval"`
}

type config struct {
	DkAgent    string  `json:"dk_agent"`
	Sender     *sender `json:"sender"`
	Service    string  `json:"service"`
	DumpSize   int     `json:"dump_size"`
	RandomDump bool    `json:"random_dump"`
	Trace      []*span `json:"trace"`
}

type tag struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type span struct {
	Resource  string        `json:"resource"`
	Operation string        `json:"operation"`
	SpanType  string        `json:"span_type"`
	Duration  time.Duration `json:"duration"`
	Error     string        `json:"error"`
	Tags      []tag         `json:"tags"`
	Children  []*span       `json:"children"`
	dumpSize  int64
}

func (sp *span) startSpanFromContext(ctx context.Context) (go2sky.Span, context.Context) {
	var tracer = go2sky.GetGlobalTracer()
	skyspan, ctx, err := tracer.CreateLocalSpan(ctx, go2sky.WithOperationName(sp.Operation))
	if err != nil {
		log.Fatalln(err.Error())
	}

	skyspan.Tag("resource.name", sp.Resource)
	skyspan.Tag("span.type", sp.SpanType)
	for _, tag := range sp.Tags {
		skyspan.Tag(go2sky.Tag(tag.Key), fmt.Sprintf("%v", tag.Value))
	}
	if len(sp.Error) != 0 {
		skyspan.Error(time.Now(), sp.Error)
	}

	if sp.dumpSize != 0 {
		buf := make([]byte, sp.dumpSize)
		rand.Read(buf)

		skyspan.Tag(go2sky.Tag("_dump_data"), hex.EncodeToString(buf))
	}

	total := int64(sp.Duration * time.Millisecond)
	d := rand.Int63n(total)
	time.Sleep(time.Duration(d))
	go func() {
		time.Sleep(time.Duration(total - d))
		skyspan.End()
	}()

	return skyspan, ctx
}

func main() {
	reporter, err := NewGRPCReporterRelay()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer reporter.Close()

	tracer, err := go2sky.NewTracer(cfg.Service, go2sky.WithReporter(reporter))
	if err != nil {
		log.Fatalln(err.Error())
	}
	go2sky.SetGlobalTracer(tracer)

	spanCount := countSpans(cfg.Trace, 0)
	log.Printf("### span count: %d\n", spanCount)
	log.Printf("### random dump: %v", cfg.RandomDump)
	if cfg.RandomDump {
		if cfg.DumpSize <= 0 {
			cfg.DumpSize = rand.Intn(924) + 100
		}
		log.Printf("### dump size: 0kb~%dkb", cfg.DumpSize)
	} else {
		log.Printf("### dump size: %dkb", cfg.DumpSize)
	}

	if cfg.RandomDump || cfg.DumpSize > 0 {
		setPerDumpSize(cfg.Trace, int64(cfg.DumpSize/spanCount)<<10, cfg.RandomDump)
	}

	root, ctx, children := startRootSpan(cfg.Trace)
	orchestrator(ctx, children)
	time.Sleep(3 * time.Second)
	root.End()

	<-globalCloser
}

func countSpans(trace []*span, c int) int {
	c += len(trace)
	for i := range trace {
		if len(trace[i].Children) != 0 {
			c = countSpans(trace[i].Children, c)
		}
	}

	return c
}

func setPerDumpSize(trace []*span, fillup int64, isRandom bool) {
	if isRandom {
		for i := range trace {
			trace[i].dumpSize = rand.Int63n(fillup)
			if len(trace[i].Children) != 0 {
				setPerDumpSize(trace[i].Children, fillup, isRandom)
			}
		}
	} else {
		for i := range trace {
			trace[i].dumpSize = fillup
			if len(trace[i].Children) != 0 {
				setPerDumpSize(trace[i].Children, fillup, isRandom)
			}
		}
	}
}

func startRootSpan(trace []*span) (root go2sky.Span, rootCtx context.Context, children []*span) {
	var sp *span
	if len(trace) == 1 {
		sp = trace[0]
		children = sp.Children
	} else {
		sp = &span{
			Operation: "startRootSpan",
			SpanType:  "web",
			Duration:  time.Duration(60 + rand.Intn(300)),
		}
		children = trace
	}
	root, rootCtx = sp.startSpanFromContext(context.Background())

	return
}

// func startRootSpan(trace []*span) (root go2sky.Span, ctx context.Context, children []*span) {
// 	var (
// 		tracer = go2sky.GetGlobalTracer()
// 		d      time.Duration
// 		err    error
// 	)
// 	if len(trace) == 1 {
// 		if root, ctx, err = tracer.CreateLocalSpan(context.TODO(), go2sky.WithOperationName(trace[0].Operation), go2sky.WithSpanType(go2sky.SpanTypeEntry)); err != nil {
// 			log.Fatalln(err.Error())
// 		}
// 		root.Tag("resource.name", trace[0].Resource)
// 		root.Tag("span.type", trace[0].SpanType)
// 		for _, tag := range trace[0].Tags {
// 			root.Tag(go2sky.Tag(tag.Key), fmt.Sprintf("%v", tag.Value))
// 		}
// 		d = trace[0].Duration * time.Millisecond
// 		children = trace[0].Children
// 		if len(trace[0].Error) != 0 {
// 			root.Error(time.Now(), trace[0].Error)
// 		}
// 	} else {
// 		if root, ctx, err = tracer.CreateLocalSpan(context.Background(), go2sky.WithOperationName("startRootSpan"), go2sky.WithSpanType(go2sky.SpanTypeEntry)); err != nil {
// 			log.Fatalln(err.Error())

// 			return
// 		}
// 		root.Tag("span.type", "web")
// 		d = time.Duration(60+rand.Intn(300)) * time.Millisecond
// 		children = trace
// 	}

// 	time.Sleep(d / 2)
// 	go func(root go2sky.Span, d time.Duration) {
// 		time.Sleep(d / 2)
// 		// root.End()
// 	}(root, d)

// 	return
// }

func orchestrator(ctx context.Context, children []*span) {
	if len(children) == 1 {
		_, ctx = children[0].startSpanFromContext(ctx)
		if len(children[0].Children) != 0 {
			orchestrator(ctx, children[0].Children)
		}
	} else {
		for k := range children {
			go func(ctx context.Context, span *span) {
				_, ctx = span.startSpanFromContext(ctx)
				if len(span.Children) != 0 {
					orchestrator(ctx, span.Children)
				}
			}(ctx, children[k])
		}
	}
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	data, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	cfg = &config{}
	if err = json.Unmarshal(data, cfg); err != nil {
		log.Fatalln(err.Error())
	}
	if cfg.Sender == nil || cfg.Sender.Threads <= 0 || cfg.Sender.SendCount <= 0 {
		log.Fatalln("invalid configuration for Sender")
	}
	if len(cfg.Trace) == 0 {
		log.Fatalln("empty trace")
	}

	rand.Seed(time.Now().UnixNano())
	agentAddress += strconv.Itoa(30000 + rand.Intn(10000))
}
