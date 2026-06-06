package grpc

import (
	"context"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"
)

type Server struct {
	identityv1.UnimplementedIdentityServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) HealthCheck(
	ctx context.Context,
	req *identityv1.HealthCheckRequest,
) (*identityv1.HealthCheckResponse, error) {

	return &identityv1.HealthCheckResponse{
		Status: "ok",
	}, nil
}