package main

import (
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/propagation"
	commonv3 "skywalking.apache.org/repo/goapi/collect/common/v3"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

type defSpan struct {
	ctx           go2sky.SegmentContext
	spanType      agentv3.SpanType
	refs          []*propagation.SpanContext
	startTime     time.Time
	endTime       time.Time
	operationName string
	peer          string
	layer         agentv3.SpanLayer
	compID        int32
	tags          []*commonv3.KeyStringValuePair
	logs          []*agentv3.Log
	isErr         bool
}

func fromReportedSpan(src go2sky.ReportedSpan) *defSpan {
	desc := &defSpan{
		ctx:           *src.Context(),
		spanType:      src.SpanType(),
		refs:          make([]*propagation.SpanContext, len(src.Refs())),
		startTime:     time.UnixMilli(src.StartTime()),
		endTime:       time.UnixMilli(src.EndTime()),
		operationName: src.OperationName(),
		peer:          src.Peer(),
		layer:         src.SpanLayer(),
		compID:        src.ComponentID(),
		tags:          src.Tags(),
		logs:          src.Logs(),
		isErr:         src.IsError(),
	}
	for i := range desc.refs {
		desc.refs[i] = &propagation.SpanContext{
			TraceID:               src.Refs()[i].TraceID,
			ParentSegmentID:       src.Refs()[i].ParentSegmentID,
			ParentService:         src.Refs()[i].ParentService,
			ParentServiceInstance: src.Refs()[i].ParentServiceInstance,
			ParentEndpoint:        src.Refs()[i].ParentEndpoint,
			AddressUsedAtClient:   src.Refs()[i].AddressUsedAtClient,
			ParentSpanID:          src.Refs()[i].ParentSpanID,
			Sample:                src.Refs()[i].Sample,
			Valid:                 src.Refs()[i].Valid,
			CorrelationContext:    src.Refs()[i].CorrelationContext,
		}
	}

	return desc
}

func (dfsp *defSpan) Context() *go2sky.SegmentContext { return &dfsp.ctx }

func (dfsp *defSpan) Refs() []*propagation.SpanContext { return dfsp.refs }

func (dfsp *defSpan) StartTime() int64 { return dfsp.startTime.UnixNano() / int64(time.Millisecond) }

func (dfsp *defSpan) EndTime() int64 { return dfsp.endTime.UnixNano() / int64(time.Millisecond) }

func (dfsp *defSpan) OperationName() string { return dfsp.operationName }

func (dfsp *defSpan) Peer() string { return dfsp.peer }

func (dfsp *defSpan) SpanType() agentv3.SpanType { return dfsp.spanType }

func (dfsp *defSpan) SpanLayer() agentv3.SpanLayer { return dfsp.layer }

func (dfsp *defSpan) IsError() bool { return dfsp.isErr }

func (dfsp *defSpan) Tags() []*commonv3.KeyStringValuePair { return dfsp.tags }

func (dfsp *defSpan) Logs() []*agentv3.Log { return dfsp.logs }

func (dfsp *defSpan) ComponentID() int32 { return dfsp.compID }
