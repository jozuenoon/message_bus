package query

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ QueryServiceServer = (*service)(nil)

func NewQueryService(repo Repository) QueryServiceServer {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo Repository
}

func (s *service) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	panic("implement me")
}

func (s *service) StreamEvents(*GetEventsRequest, QueryService_StreamEventsServer) error {
	return status.Error(codes.Unimplemented, "sorry we are work in progress")
}

func (s *service) GetBatteryVoltage(context.Context, *GetBatteryVoltageRequest) (*GetBatteryVoltageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "sorry we are work in progress")
}
