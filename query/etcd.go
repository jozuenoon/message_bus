package query

import (
	"context"
	"time"

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

func (r *etcdRepository) GetEvents(ctx context.Context, detectors []string, after, before time.Time, limit int64) ([]*Event, error) {
	panic("implement me")
}
