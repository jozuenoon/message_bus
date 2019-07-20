package pkg

import (
	"context"

	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/healthcheck"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"net"
)

func HealthCheck(_ context.Context, g *run.Group, port string, logger log15.Logger) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	g.Add(func() error {
		grpcServer := grpc.NewServer()
		svc := healthcheck.New(logger)
		grpc_health_v1.RegisterHealthServer(grpcServer, svc)
		return grpcServer.Serve(ln)
	}, func(error) {
		ln.Close()
	})
	return nil
}
