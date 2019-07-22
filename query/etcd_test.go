// +build integration

package query

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jozuenoon/message_bus/collector"

	"github.com/coreos/etcd/clientv3"
	"github.com/google/uuid"

	"github.com/inconshreveable/log15"
)

func init() {
	if v, ok := os.LookupEnv("ETCD_ENDPOINTS"); ok {
		etcdEndpoints = strings.Split(v, ",")
	}
}

var etcdEndpoints = []string{"http://etcd:2379"}

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
			if got := etcdEventKeys(tt.namespace, tt.detectors); !reflect.DeepEqual(got, tt.want) {
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

type tval struct {
	Timestamp int64
}

func Test_etcdRepository_GetEvents(t *testing.T) {
	prefix := "/integration_tests"
	repo, err := NewEtcdRepository(prefix, etcdEndpoints, log15.New())
	if err != nil {
		t.Fatalf("failed to create etcd repository: %v", err)
	}
	r := repo.(*etcdRepository)
	tests := []struct {
		name      string
		seed      string
		detectors []string
		after     time.Time
		before    time.Time
		limit     int64
		want      []*Event
		wantErr   bool
	}{
		{
			name:      "basic",
			seed:      uuid.New().String(),
			detectors: []string{"xxx-1"},
			limit:     100,
			want: []*Event{
				{
					DetectorID: "xxx-1",
					DeviceID:   "ddd-1",
					Time:       time.Unix(0, 3000),
				},
				{
					DetectorID: "xxx-1",
					DeviceID:   "ddd-2",
					Time:       time.Unix(0, 2000),
				},
				{
					DetectorID: "xxx-1",
					DeviceID:   "ddd-3",
					Time:       time.Unix(0, 1000),
				},
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.seed)
			r.prefix = tt.seed
			for _, ev := range tt.want {
				putETCDEvent(t, tt.seed, ev.DetectorID, ev.DeviceID, ev.Time, r.cli)
			}
			got, err := r.GetEvents(context.Background(), tt.detectors, tt.after, tt.before, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func putETCDEvent(t *testing.T, namespace, detectorID, deviceID string, ts time.Time, cli *clientv3.Client) {
	t.Helper()
	key := collector.EventKey(namespace, detectorID, deviceID, ts)
	bval, err := json.Marshal(tval{
		Timestamp: ts.UnixNano(),
	})
	if err != nil {
		t.Fatalf("failed to marshal json")
	}
	val := base64.StdEncoding.EncodeToString(bval)
	_, err = cli.Put(context.Background(), key, val)
	if err != nil {
		t.Fatalf("failed to check key: %v", err)
	}
}
