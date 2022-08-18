package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	confv3 "skywalking.apache.org/repo/goapi/collect/agent/configuration/v3"
	commv3 "skywalking.apache.org/repo/goapi/collect/common/v3"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	mgmtv3 "skywalking.apache.org/repo/goapi/collect/management/v3"
)

func startAgent() {
	log.Printf("### start Skywalking agent %s\n", agentAddress)

	listener, err := net.Listen("tcp", agentAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gsrv := grpc.NewServer()
	agentv3.RegisterTraceSegmentReportServiceServer(gsrv, &TraceSegmentReportServer{})
	mgmtv3.RegisterManagementServiceServer(gsrv, &ManagementServerV3{})
	confv3.RegisterConfigurationDiscoveryServiceServer(gsrv, &DiscoveryServerV3{})

	if err := gsrv.Serve(listener); err != nil {
		log.Println(err.Error())
	}
}

type TraceSegmentReportServer struct {
	agentv3.UnimplementedTraceSegmentReportServiceServer
}

func (*TraceSegmentReportServer) Collect(srv agentv3.TraceSegmentReportService_CollectServer) error {
	for {
		select {
		case <-globalCloser:
			break
		default:
		}

		segobj, err := srv.Recv()
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, io.EOF) {
				return srv.SendAndClose(&commv3.Commands{})
			}
			continue
		}

		log.Println(*segobj)
	}
}

func (*TraceSegmentReportServer) CollectInSync(ctx context.Context, coll *agentv3.SegmentCollection) (*commv3.Commands, error) {
	return &commv3.Commands{}, nil
}

type ManagementServerV3 struct {
	mgmtv3.UnimplementedManagementServiceServer
}

func (*ManagementServerV3) ReportInstanceProperties(ctx context.Context, mgmt *mgmtv3.InstanceProperties) (*commv3.Commands, error) {
	return &commv3.Commands{}, nil
}

func (*ManagementServerV3) KeepAlive(ctx context.Context, ping *mgmtv3.InstancePingPkg) (*commv3.Commands, error) {
	return &commv3.Commands{}, nil
}

type DiscoveryServerV3 struct {
	confv3.UnimplementedConfigurationDiscoveryServiceServer
}

func (*DiscoveryServerV3) FetchConfigurations(ctx context.Context, cfgReq *confv3.ConfigurationSyncRequest) (*commv3.Commands, error) {
	return &commv3.Commands{}, nil
}
