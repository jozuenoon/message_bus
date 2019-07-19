package pkg

import (
	"context"
	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/collector"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"net"
)

func ETCDCollectorServiceServer(_ context.Context, g *run.Group, prefix string, port string, endpoints []string, logger log15.Logger) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	repo, err := collector.NewEtcdRepository(prefix, endpoints, logger)
	if err != nil {
		return err
	}
	svc := collector.NewCollectorService(repo)
	g.Add(func() error {
		grpcServer := grpc.NewServer()
		collector.RegisterCollectorServiceServer(grpcServer, svc)
		return grpcServer.Serve(ln)
	}, func(error) {
		ln.Close()
	}) // TODO: shutdown ETCD?
	return nil
}