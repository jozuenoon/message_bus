package collector

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"path"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/pkg/types"
	"github.com/oklog/ulid/v2"
)

const detectorsPrefix = "detectors"
const linkPrefix = "links"

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

func (r *etcdRepository) CreateDetectorLink(ctx context.Context, destDetectorID, srcDetectorID string, maxSeconds int64) error {
	key := detectorLinkKey(r.prefix, destDetectorID, srcDetectorID)
	value := struct {
		MaxSeconds int64
	}{
		MaxSeconds: maxSeconds,
	}
	return r.save(ctx, key, value)
}

func (r *etcdRepository) CreateDetectionEvent(ctx context.Context, detectorID,
	deviceID string, timestamp time.Time) error {

	key := eventKey(r.prefix, detectorID, deviceID, timestamp)
	value := struct {
		Timestamp int64
	}{
		Timestamp: timestamp.UnixNano(),
	}
	return r.save(ctx, key, value)
}

func (r *etcdRepository) CreateActiveDetectors(ctx context.Context, detectorID,
	activeDetectors string, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateBatteryVoltage(ctx context.Context, detectorID string,
	voltage float64, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateDetectionCount(ctx context.Context, detectorID,
	detectorType string, count int64, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateCoordinates(ctx context.Context, detectorID string,
	latitude, longitude types.DecimalDegrees) error {
	panic("implement me")
}

func (r *etcdRepository) save(ctx context.Context, key string, value interface{}) error {
	bval, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to serialize detector link: %v", err)
	}
	val := base64.StdEncoding.EncodeToString(bval)
	_, err = r.cli.Put(ctx, key, val)
	if err != nil {
		return fmt.Errorf("failed to create detector link: %v", err)
	}
	return nil
}

func detectorLinkKey(namespace, destDetectorID, srcDetectorID string) string {
	return "/" + path.Join(namespace, linkPrefix, destDetectorID, srcDetectorID)
}

func eventKey(namespace, detectorID, deviceID string, timestamp time.Time) string {
	ts := timestamp.Unix()
	timestamp.UnixNano()
	timeKey := strconv.FormatInt(ts-ts%100, 10)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(timestamp.UnixNano())), 0)
	u := ulid.MustNew(ulid.Timestamp(timestamp), entropy)
	return "/" + path.Join(namespace, detectorsPrefix, detectorID, timeKey, u.String()+"."+deviceID)
}
