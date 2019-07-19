package collector

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ CollectorServiceServer = (*service)(nil)

func NewCollectorService(repo Repository) CollectorServiceServer {
	return &service{
		repo: repo,
	}
}


type service struct {
	repo Repository
}

func (s *service) CreateEventLog(context.Context, *EventLog) (*empty.Empty, error) {
	panic("implement me")
}

func (s *service) StreamEventLog(CollectorService_StreamEventLogServer) error {
	panic("implement me")
}

func (s *service) CreateCoordinates(context.Context, *Coordinates) (*empty.Empty, error) {
	panic("implement me")
}

func (s *service) CreateDetectorStatus(context.Context, *DetectorStatus) (*empty.Empty, error) {
	panic("implement me")
}




