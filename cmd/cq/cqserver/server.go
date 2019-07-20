package cqserver

import (
	"context"
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/cmd/pkg"
	"github.com/oklog/run"
)

func New(collectorPort, queryPort, healthCheckPort, etcdPrefix string, etcdEndpoints []string, logger log15.Logger) *server {
	return &server{
		collectorPort: collectorPort,
		queryPort: queryPort,
		healthCheckPort: healthCheckPort,
		etcdPrefix: etcdPrefix,
		etcdEndpoints: etcdEndpoints,
		logger: logger,
	}
}


type server struct {
	collectorPort string
	queryPort string
	healthCheckPort string
	etcdPrefix string
	etcdEndpoints []string
	logger log15.Logger
}

func (s *server) Run(ctx context.Context) error {
	g := &run.Group{}
	cctx, cancel := context.WithCancel(ctx)

	// If context is canceled deeper in application this will terminate whole server.
	g.Add(func () error {
		<-cctx.Done()
		return fmt.Errorf("context canceled")
	}, func (error) {
		cancel()
	})

	// Add signal handler
	pkg.SignalHandler(ctx, g, s.logger.New("module", "signal_handler"))

	err := pkg.ETCDCollectorServiceServer(ctx, g, s.etcdPrefix, s.collectorPort, s.etcdEndpoints, s.logger.New("service", "collector"))
	if err != nil {
		return err
	}
	err = pkg.ETCDQueryServiceServer(ctx, g, s.etcdPrefix, s.queryPort, s.etcdEndpoints, s.logger.New("service", "query"))
	if err != nil {
		return err
	}
	err = pkg.HealthCheck(ctx, g, s.healthCheckPort, s.logger.New("service", "health_check"))
	if err != nil {
		return err
	}
	return g.Run()
}