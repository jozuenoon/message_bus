// +build integration

package collector

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/coreos/etcd/clientv3"
	"github.com/inconshreveable/log15"
)

func init() {
	if v, ok := os.LookupEnv("ETCD_ENDPOINTS"); ok {
		etcdEndpoints = strings.Split(v, ",")
	}
}

var etcdEndpoints = []string{"http://etcd:2379"}

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
			got := EventKey(tt.namespace, tt.detectorID, tt.deviceID, tt.timestamp)
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

func Test_etcdRepository_CreateDetectionEvent(t *testing.T) {
	type tval struct {
		Timestamp int64
	}
	prefix := "/integration_tests"
	repo, err := NewEtcdRepository(prefix, etcdEndpoints, log15.New())
	if err != nil {
		t.Fatalf("failed to create etcd repository: %v", err)
	}
	r := repo.(*etcdRepository)
	tests := []struct {
		name       string
		detectorID string
		deviceID   string
		timestamp  time.Time
		wantErr    bool
	}{
		{
			name:       "basic",
			detectorID: "xxx-1",
			deviceID:   "ddd-1",
			timestamp:  time.Unix(0, 1000),
		},
		{
			name:       "basic",
			detectorID: "xxx-2",
			deviceID:   "ddd-2",
			timestamp:  time.Unix(0, 2000),
		},
		{
			name:       "basic",
			detectorID: "xxx-3",
			deviceID:   "ddd-3",
			timestamp:  time.Unix(0, 3000),
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			if err := r.CreateDetectionEvent(context.Background(), tt.detectorID, tt.deviceID, tt.timestamp); (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.CreateDetectionEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			value := &tval{
				Timestamp: tt.timestamp.UnixNano(),
			}
			assertETCDKeyExists(t, EventKey(prefix, tt.detectorID, tt.deviceID, tt.timestamp), value, &tval{}, r.cli)
		})
	}
}

func Test_etcdRepository_save(t *testing.T) {
	type tval struct {
		Timestamp string
	}
	prefix := "/integration_tests"
	repo, err := NewEtcdRepository(prefix, etcdEndpoints, log15.New())
	if err != nil {
		t.Fatalf("failed to create etcd repository: %v", err)
	}
	r := repo.(*etcdRepository)
	tests := []struct {
		name    string
		key     string
		value   interface{}
		wantErr bool
	}{
		{
			name:    "push",
			key:     "example-key-1",
			value:   &tval{time.Now().Format(time.RFC3339)},
			wantErr: false,
		},
		{
			name:    "push",
			key:     "example-key-2",
			value:   &tval{time.Now().Format(time.RFC3339)},
			wantErr: false,
		},
		{
			name:    "push",
			key:     "example-key-3",
			value:   &tval{time.Now().Format(time.RFC3339)},
			wantErr: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			if err := r.save(context.Background(), tt.key, tt.value); (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.save() error = %v, wantErr %v", err, tt.wantErr)
			}
			assertETCDKeyExists(t, tt.key, tt.value, &tval{}, r.cli)
		})
	}
}

func assertETCDKeyExists(t *testing.T, key string, value, result interface{}, cli *clientv3.Client) {
	t.Helper()
	resp, err := cli.Get(context.Background(), key)
	if err != nil {
		t.Fatalf("failed to check key: %v", err)
	}
	if resp.Count != 1 {
		t.Fatalf("have number of returned values not equal 1: %v", resp.Kvs)
	}
	rawValue := resp.Kvs[0].Value
	val, err := base64.StdEncoding.DecodeString(string(rawValue))
	if err != nil {
		t.Fatalf("failed to decode base64: %v", err)
	}
	err = json.Unmarshal(val, result)
	if err != nil {
		fmt.Println(string(val))
		t.Fatalf("failed to unmarshal value: %s", string(val))
	}
	assert.Equal(t, value, result)
}

func Test_etcdRepository_CreateDetectorLink(t *testing.T) {
	r := &etcdRepository{}
	assert.Panics(t, func() { _ = r.CreateDetectorLink(context.Background(), "a", "b", 5) })
}

func Test_etcdRepository_CreateActiveDetectors(t *testing.T) {
	r := &etcdRepository{}
	assert.Panics(t, func() { _ = r.CreateActiveDetectors(context.Background(), "a", "b", time.Time{}) })
}

func Test_etcdRepository_CreateBatteryVoltage(t *testing.T) {
	r := &etcdRepository{}
	assert.Panics(t, func() { _ = r.CreateBatteryVoltage(context.Background(), "a", 0, time.Time{}) })
}

func Test_etcdRepository_CreateDetectionCount(t *testing.T) {
	r := &etcdRepository{}
	assert.Panics(t, func() { _ = r.CreateDetectionCount(context.Background(), "a", "b", 1, time.Time{}) })
}

func Test_etcdRepository_CreateCoordinates(t *testing.T) {
	r := &etcdRepository{}
	assert.Panics(t, func() { _ = r.CreateCoordinates(context.Background(), "a", 1, 1) })
}
