package collector

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/collector/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_service_CreateEventLog(t *testing.T) {
	tests := []struct {
		name     string
		eventLog *EventLog
		wantErr  bool
	}{
		{
			name: "simple",
			eventLog: &EventLog{
				Loc: &Coordinates{
					DetectorId: "my-detector-id",
				},
				Events: []*DetectionEvent{
					{
						DeviceId: &DetectionEvent_Bluetooth{
							Bluetooth: "my-bt-device-id",
						},
						Time: []*timestamp.Timestamp{
							{Seconds: 100, Nanos: 200},
							{Seconds: 200, Nanos: 300},
						},
					},
					{
						DeviceId: &DetectionEvent_Wifi{
							Wifi: "my-wifi-device-id",
						},
						Time: []*timestamp.Timestamp{
							{Seconds: 100, Nanos: 200},
							{Seconds: 200, Nanos: 300},
						},
					},
					{
						DeviceId: &DetectionEvent_Mobile{
							Mobile: "my-mobile-device-id",
						},
						Time: []*timestamp.Timestamp{
							{Seconds: 100, Nanos: 200},
							{Seconds: 200, Nanos: 300},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := log15.New("test", "collector")
			repoMock := &mocks.Repository{}
			mockResultFn := func(detectorID string, deviceID string, timestamp time.Time) error {
				// Confirm existence of arguments in original event.
				if detectorID != tt.eventLog.Loc.DetectorId {
					return fmt.Errorf("wrong value of detector ID want=%v, got=%v", tt.eventLog.Loc.DetectorId, detectorID)
				}
				for _, ev := range tt.eventLog.Events {
					evdID, err := decodeDeviceID(ev)
					if err != nil {
						return err
					}
					if evdID == deviceID {
						for _, ts := range ev.Time {
							t := time.Unix(ts.Seconds, int64(ts.Nanos))
							if t == timestamp {
								return nil
							}
						}
					}
				}
				return fmt.Errorf("failed to find device detection event")
			}
			repoMock.On("CreateDetectionEvent",
				mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).
				Return(mockResultFn, nil)

			s := NewCollectorService(repoMock, logger)
			_, err := s.CreateEventLog(context.Background(), tt.eventLog)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateEventLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_service_CreateDetectorLink(t *testing.T) {
	tests := []struct {
		name    string
		dLink   *DetectorLink
		wantErr bool
	}{
		{
			name: "simple",
			dLink: &DetectorLink{
				DestDetectorId: "my-dest-detector-id",
				SrcDetectors: []*DetectorLink_SourceDetector{
					{
						DetectorId: "src-detector-id-1",
						MaxSeconds: 1000,
					},
					{
						DetectorId: "src-detector-id-2",
						MaxSeconds: 600,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := log15.New("test", "collector")
			repoMock := &mocks.Repository{}
			mockResultFn := func(destDetectorID, srcDetectorID string, maxSeconds int64) error {
				// Confirm existence of arguments in original request.
				if destDetectorID != tt.dLink.DestDetectorId {
					return fmt.Errorf("wrong value of detector ID want=%v, got=%v", tt.dLink.DestDetectorId, destDetectorID)
				}
				for _, srcdID := range tt.dLink.SrcDetectors {
					if srcdID.DetectorId == srcDetectorID && srcdID.MaxSeconds == maxSeconds {
						return nil
					}
				}
				return fmt.Errorf("failed to find device link")
			}

			repoMock.On("CreateDetectorLink",
				mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int64")).
				Return(mockResultFn, nil)

			s := NewCollectorService(repoMock, logger)
			_, err := s.CreateDetectorLink(context.Background(), tt.dLink)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateDetectorLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
