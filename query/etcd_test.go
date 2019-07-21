package query

import (
	"reflect"
	"testing"
	"time"
)

func Test_keys(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		detectors []string
		want      []string
	}{
		{
			name:      "basic",
			namespace: "ns",
			detectors: []string{"xxx-1", "xxx-2"},
			want:      []string{"/ns/detectors/xxx-1", "/ns/detectors/xxx-2"},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			if got := keys(tt.namespace, tt.detectors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseKey(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want *Event
	}{
		{
			name: "basic",
			key:  "/ns/detectors/xxx-1/94994400/some-ulid.my-device-id",
			want: &Event{
				DetectorID: "xxx-1",
				DeviceID:   "my-device-id",
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			if got := parseKey(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventValue(t *testing.T) {
	tests := []struct {
		name     string
		rawValue []byte
		want     time.Time
		wantErr  bool
	}{
		{
			name:     "basic",
			rawValue: []byte("eyJUaW1lc3RhbXAiOjE1NjM2OTE4MDAxMzk3Nzg0MzZ9"),
			want:     time.Unix(0, 1563691800139778436),
			wantErr:  false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			got, err := eventValue(tt.rawValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("eventValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
