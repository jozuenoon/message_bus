package collector

import (
	"context"
	"fmt"
	"time"

	"github.com/inconshreveable/log15"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
)

var _ CollectorServiceServer = (*service)(nil)

func NewCollectorService(repo Repository, logger log15.Logger) CollectorServiceServer {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

type service struct {
	repo   Repository
	logger log15.Logger
}

func (s *service) CreateEventLog(_ context.Context, eventLog *EventLog) (*empty.Empty, error) {
	detectorID := eventLog.Loc.DetectorId
	for _, ev := range eventLog.Events {
		for _, ts := range ev.Time {
			t := time.Unix(ts.Seconds, int64(ts.Nanos))

			deviceID, err := decodeDeviceID(ev)
			if err != nil {
				s.logger.Error("failed to create event", "err", err)
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}

			err = s.repo.CreateDetectionEvent(detectorID, deviceID, t)
			if err != nil {
				s.logger.Error("failed to create event", "err", err)
				return nil, status.Error(codes.Internal, "failed to save detection event")
			}
		}
	}
	return nil, nil
}

func (s *service) CreateDetectorLink(_ context.Context, dLink *DetectorLink) (*empty.Empty, error) {
	for _, srcDetector := range dLink.SrcDetectors {
		err := s.repo.CreateDetectorLink(dLink.DestDetectorId, srcDetector.DetectorId, srcDetector.MaxSeconds)
		if err != nil {
			s.logger.Error("failed to save detector link", "err", err)
			return nil, status.Error(codes.Internal, "failed to save detector link")
		}
	}
	return nil, nil
}

func (s *service) StreamEventLog(CollectorService_StreamEventLogServer) error {
	return status.Error(codes.Unimplemented, "sorry we are work in progress")
}

func (s *service) CreateCoordinates(context.Context, *Coordinates) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "sorry we are work in progress")
}

func (s *service) CreateDetectorStatus(context.Context, *DetectorStatus) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "sorry we are work in progress")
}

func decodeDeviceID(ev *DetectionEvent) (string, error) {
	var deviceID string
	switch tev := ev.DeviceId.(type) {
	case *DetectionEvent_Wifi:
		deviceID = tev.Wifi
	case *DetectionEvent_Bluetooth:
		deviceID = tev.Bluetooth
	case *DetectionEvent_Mobile:
		deviceID = tev.Mobile
	default:
		return "", fmt.Errorf("failed to decode detection event")
	}
	return deviceID, nil
}
