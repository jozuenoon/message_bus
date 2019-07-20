package query

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/inconshreveable/log15"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ QueryServiceServer = (*service)(nil)

func NewQueryService(repo Repository, logger log15.Logger) QueryServiceServer {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

type service struct {
	repo   Repository
	logger log15.Logger
}

func (s *service) GetEvents(ctx context.Context, req *GetEventsRequest) (*GetEventsResponse, error) {
	after := time.Unix(req.Range.After.Seconds, int64(req.Range.After.Nanos))
	before := time.Unix(req.Range.Before.Seconds, int64(req.Range.Before.Nanos))
	events, err := s.repo.GetEvents(ctx, req.DetectorIds, after, before, req.Limit)
	if err != nil {
		s.logger.Crit("failed to get events", "err", err)
		return nil, status.Error(codes.Internal, "failed to get events")
	}

	mapper := make(map[string]*DetectionEvent)
	for _, ev := range events {
		key := ev.DetectorID + ev.DeviceID
		if _, ok := mapper[key]; !ok {
			mapper[key] = &DetectionEvent{
				DeviceId:   ev.DeviceID,
				DetectorId: ev.DetectorID,
				Time:       []*timestamp.Timestamp{{Seconds: ev.Time.Unix(), Nanos: int32(ev.Time.Nanosecond())}},
			}
			continue
		}

		devent := mapper[key]
		ts := &timestamp.Timestamp{Seconds: ev.Time.Unix(), Nanos: int32(ev.Time.Nanosecond())}
		devent.Time = append(devent.Time, ts)
	}

	detectionEvents := make([]*DetectionEvent, 0, len(mapper))
	for _, ev := range mapper {
		detectionEvents = append(detectionEvents, ev)
	}

	return &GetEventsResponse{
		Events: detectionEvents,
	}, nil
}

func (s *service) StreamEvents(req *GetEventsRequest, stream QueryService_StreamEventsServer) error {
	return status.Error(codes.Unimplemented, "sorry we are work in progress")
}

func (s *service) GetBatteryVoltage(context.Context, *GetBatteryVoltageRequest) (*GetBatteryVoltageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "sorry we are work in progress")
}
