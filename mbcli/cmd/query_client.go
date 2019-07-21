package cmd

import (
	"github.com/jozuenoon/message_bus/query"
	"google.golang.org/grpc"
)

func NewQueryClient(host string) (query.QueryServiceClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return query.NewQueryServiceClient(conn), nil
}
