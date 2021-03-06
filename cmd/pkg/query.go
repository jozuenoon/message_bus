package pkg

import (
	"context"
	"net"

	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/query"
	"github.com/oklog/run"
	"google.golang.org/grpc"
)

func ETCDQueryServiceServer(_ context.Context, g *run.Group, prefix, port string, endpoints []string, logger log15.Logger) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	repo, err := query.NewEtcdRepository(prefix, endpoints, logger)
	if err != nil {
		return err
	}
	svc := query.NewQueryService(repo, logger)
	g.Add(func() error {
		grpcServer := grpc.NewServer()
		query.RegisterQueryServiceServer(grpcServer, svc)
		return grpcServer.Serve(ln)
	}, func(error) {
		ln.Close()
	})
	return nil
}
