package query_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/jozuenoon/message_bus/query"

	"github.com/stretchr/testify/assert"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/jozuenoon/message_bus/query/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/inconshreveable/log15"
)

func Test_service_GetEvents(t *testing.T) {
	tests := []struct {
		name    string
		req     *query.GetEventsRequest
		want    *query.GetEventsResponse
		wantErr bool
	}{
		{
			name: "basic",
			req: &query.GetEventsRequest{
				DetectorIds: []string{"xxx-1"},
			},
			want: &query.GetEventsResponse{
				Events: []*query.DetectionEvent{
					{
						DetectorId: "xxx-1",
						DeviceId:   "my-device",
						Time: []*timestamp.Timestamp{
							{
								Seconds: 100,
							},
							{
								Seconds: 200,
							},
						},
					},
				},
			},
		},
		{
			name: "with time",
			req: &query.GetEventsRequest{
				DetectorIds: []string{"xxx-1"},
				Range: &query.Range{
					After:  &timestamp.Timestamp{Seconds: 100},
					Before: &timestamp.Timestamp{Seconds: 200},
				},
			},
			want: &query.GetEventsResponse{
				Events: []*query.DetectionEvent{
					{
						DetectorId: "xxx-2",
						DeviceId:   "my-device-3",
						Time: []*timestamp.Timestamp{
							{
								Seconds: 120,
							},
							{
								Seconds: 130,
							},
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			logger := log15.New("test", "collector")
			repoMock := &mocks.Repository{}
			mockResultFn := func(ctx context.Context, detectors []string, after, before time.Time, limit int64) []*query.Event {
				// Input expectations.
				assert.ElementsMatch(t, tt.req.DetectorIds, detectors, "detectors list does not match")
				if tt.req.Range != nil {
					assert.Equal(t, timestampToTime(tt.req.Range.After), after, "after does not match")
					assert.Equal(t, timestampToTime(tt.req.Range.Before), before, "before does not match")
				} else {
					assert.Equal(t, time.Time{}, after, "after does not match")
					assert.Equal(t, time.Time{}, before, "before does not match")
				}
				assert.Equal(t, tt.req.Limit, limit)

				// Denormalize.
				events := make([]*query.Event, 0)
				for _, ev := range tt.want.Events {
					for _, dev := range ev.Time {
						events = append(events, &query.Event{DetectorID: ev.DetectorId, DeviceID: ev.DeviceId, Time: timestampToTime(dev)})
					}
				}
				return events
			}
			repoMock.On("GetEvents",
				mock.Anything, mock.AnythingOfType("[]string"), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time"), mock.AnythingOfType("int64")).
				Return(mockResultFn, nil)

			s := query.NewQueryService(repoMock, logger)

			got, err := s.GetEvents(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func timestampToTime(in *timestamp.Timestamp) time.Time {
	return time.Unix(in.Seconds, int64(in.Nanos))
}
