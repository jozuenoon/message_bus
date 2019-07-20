package collector

import (
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/pkg/types"
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

func (r *etcdRepository) CreateDetectionEvent(detectorID string, deviceID string, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateActiveDetectors(detectorID string, activeDetectors string, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateBatteryVoltage(detectorID string, voltage float64, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateDetectionCount(detectorID string, detectorType string, count int64, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateCoordinates(detectorID string, latitude, longitude types.DecimalDegrees) error {
	panic("implement me")
}
