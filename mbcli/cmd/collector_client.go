package cmd

import (
	"github.com/jozuenoon/message_bus/collector"
	"google.golang.org/grpc"
)

func NewCollectorClient(host string) (collector.CollectorServiceClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return collector.NewCollectorServiceClient(conn), nil
}
