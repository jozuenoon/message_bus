package cmd

import (
	"context"

	"google.golang.org/grpc"

	"github.com/jozuenoon/message_bus/query"
)

var _ query.QueryServiceClient = (*queryClient)(nil)

type queryClient struct{}

func (q *queryClient) GetEvents(ctx context.Context, in *query.GetEventsRequest, opts ...grpc.CallOption) (*query.GetEventsResponse, error) {
	panic("implement me")
}

func (q *queryClient) StreamEvents(ctx context.Context, in *query.GetEventsRequest, opts ...grpc.CallOption) (query.QueryService_StreamEventsClient, error) {
	panic("implement me")
}

func (q *queryClient) GetBatteryVoltage(ctx context.Context, in *query.GetBatteryVoltageRequest, opts ...grpc.CallOption) (*query.GetBatteryVoltageResponse, error) {
	panic("implement me")
}
