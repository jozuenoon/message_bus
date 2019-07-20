package healthcheck

import (
	"github.com/inconshreveable/log15"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

var _ grpc_health_v1.HealthServer = (*server)(nil)

func New(logger log15.Logger) grpc_health_v1.HealthServer {
	return &server{
		logger: logger,
	}
}


type server struct {
	logger log15.Logger
}

func (s *server) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	// TODO - implement, for now return ok.
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *server) Watch(_ *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	// TODO - implement, for now return unimplemented.
	return status.Error(codes.Unimplemented, "watch is not implemented")
}
