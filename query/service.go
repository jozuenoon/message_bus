package query

import "context"

var _ QueryServiceServer = (*service)(nil)

type service struct {}

func (s *service) StreamEvents(*GetEventsRequest, QueryService_StreamEventsServer) error {
	panic("implement me")
}

func (s *service) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	panic("implement me")
}


func (s *service) GetBatteryVoltage(context.Context, *GetBatteryVoltageRequest) (*GetBatteryVoltageResponse, error) {
	panic("implement me")
}