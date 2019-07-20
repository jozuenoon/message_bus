package cmd

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jozuenoon/message_bus/collector"
	"google.golang.org/grpc"
)

var _ collector.CollectorServiceClient = (*collectorClient)(nil)

type collectorClient struct{}

func (c *collectorClient) CreateEventLog(ctx context.Context, in *collector.EventLog, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (c *collectorClient) StreamEventLog(ctx context.Context, opts ...grpc.CallOption) (collector.CollectorService_StreamEventLogClient, error) {
	panic("implement me")
}

func (c *collectorClient) CreateCoordinates(ctx context.Context, in *collector.Coordinates, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (c *collectorClient) CreateDetectorStatus(ctx context.Context, in *collector.DetectorStatus, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}
