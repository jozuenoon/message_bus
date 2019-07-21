package query

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/jozuenoon/message_bus/collector"

	"github.com/coreos/etcd/clientv3"
	"github.com/inconshreveable/log15"
)

var _ Repository = (*etcdRepository)(nil)

func NewEtcdRepository(prefix string, endpoints []string, logger log15.Logger) (Repository, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &etcdRepository{
		prefix: prefix,
		logger: logger,
		cli:    cli,
	}, nil
}

type etcdRepository struct {
	prefix string
	logger log15.Logger
	cli    *clientv3.Client
}

// GetEvents retrieves events with filters applied.
func (r *etcdRepository) GetEvents(ctx context.Context, detectors []string, after, before time.Time, limit int64) ([]*Event, error) {
	// TODO: supports only detector selection, other filers are ignored.
	k := keys(r.prefix, detectors)
	options := append(clientv3.WithLastKey(), clientv3.WithPrefix(), clientv3.WithLimit(limit))
	events := []*Event{}
	for _, key := range k {
		resp, err := r.cli.Get(ctx, key, options...)
		if err != nil {
			return nil, fmt.Errorf("failed to get keys from ETCD: %v", err)
		}
		for _, kv := range resp.Kvs {
			ev := parseKey(string(kv.Key))
			ev.Time, err = eventValue(kv.Value)
			if err != nil {
				return nil, err
			}
			events = append(events, ev)
		}
	}
	return events, nil
}

// TODO: ETCD key operations certainly deserves library on it's own to keep things consistent across services.
func keys(namespace string, detectors []string) []string {
	keys := make([]string, 0, len(detectors))
	for _, detectorID := range detectors {
		keys = append(keys, "/"+path.Join(namespace, collector.ETCDDetectorsPrefix, detectorID))
	}
	return keys
}

func eventValue(rawValue []byte) (time.Time, error) {
	tsj := struct {
		Timestamp int64
	}{}
	vdec, err := base64.StdEncoding.DecodeString(string(rawValue))
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to base64 decode value: %v", err)
	}
	err = json.Unmarshal(vdec, &tsj)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to unmarshal value: %v", err)
	}
	return time.Unix(0, tsj.Timestamp), nil
}

// parseKey returns detectorID and deviceID from key
func parseKey(key string) *Event {
	sp := strings.Split(key, "/")
	// DetectorID should be 3 element since key starts with slash.
	detectorID := sp[3]

	// Lookups event keys and detaches ULID.
	deviceID := strings.Split(sp[len(sp)-1], ".")[1]
	return &Event{
		DeviceID:   deviceID,
		DetectorID: detectorID,
	}
}
