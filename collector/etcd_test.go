package collector

import (
	"strings"
	"testing"
	"time"
)

func Test_detectorLinkKey(t *testing.T) {
	tests := []struct {
		name           string
		namespace      string
		destDetectorID string
		srcDetectorID  string
		want           string
	}{
		{
			name:           "simple",
			namespace:      "somenamespace",
			destDetectorID: "dest-det-id",
			srcDetectorID:  "src-det-id",
			want:           "/somenamespace/links/dest-det-id/src-det-id",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(test.name, func(t *testing.T) {
			if got := detectorLinkKey(tt.namespace, tt.destDetectorID, tt.srcDetectorID); got != tt.want {
				t.Errorf("detectorLinkPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventKey(t *testing.T) {
	tests := []struct {
		name       string
		namespace  string
		detectorID string
		deviceID   string
		timestamp  time.Time
		want       string
	}{
		{
			name:       "simple",
			namespace:  "ns",
			detectorID: "detID",
			deviceID:   "devID",
			timestamp:  time.Unix(1563568110, 0),
			want:       "/ns/detectors/detID/1563568100/devID",
		},
		{
			name:       "simple",
			namespace:  "ns",
			detectorID: "detID",
			deviceID:   "devID",
			timestamp:  time.Unix(1563568120, 0),
			want:       "/ns/detectors/detID/1563568100/devID",
		},
		{
			name:       "simple",
			namespace:  "ns",
			detectorID: "detID",
			deviceID:   "devID",
			timestamp:  time.Unix(1563568180, 0),
			want:       "/ns/detectors/detID/1563568100/devID",
		},
		{
			name:       "simple",
			namespace:  "ns",
			detectorID: "detID",
			deviceID:   "devID",
			timestamp:  time.Unix(1563568190, 0),
			want:       "/ns/detectors/detID/1563568100/devID",
		},
		{
			name:       "simple",
			namespace:  "ns",
			detectorID: "detID",
			deviceID:   "devID",
			timestamp:  time.Unix(1563568200, 0),
			want:       "/ns/detectors/detID/1563568200/devID",
		},
		{
			name:       "simple",
			namespace:  "ns",
			detectorID: "detID",
			deviceID:   "devID",
			timestamp:  time.Unix(1563568210, 0),
			want:       "/ns/detectors/detID/1563568200/devID",
		},
		{
			name:       "simple",
			namespace:  "ns",
			detectorID: "detID",
			deviceID:   "devID",
			timestamp:  time.Unix(1563568220, 0),
			want:       "/ns/detectors/detID/1563568200/devID",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			got := eventKey(tt.namespace, tt.detectorID, tt.deviceID, tt.timestamp)
			// Need to remove non deterministic part of ulid.
			sp := strings.Split(got, "/")
			last := strings.Split(sp[len(sp)-1], ".")
			sp[len(sp)-1] = last[1]
			got = strings.Join(sp, "/")
			if got != tt.want {
				t.Errorf("detectorEventKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
