package grpc

import (
	"context"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/service"
)

type Server struct {
	identityv1.UnimplementedIdentityServiceServer

	identityService *service.IdentityService
}

func NewServer(
	identityService *service.IdentityService,
) *Server {
	return &Server{
		identityService: identityService,
	}
}

func (s *Server) HealthCheck(
	ctx context.Context,
	req *identityv1.HealthCheckRequest,
) (*identityv1.HealthCheckResponse, error) {
	res, err := s.identityService.HealthCheck(ctx)
	if err != nil {
		return nil, err
	}
	return &identityv1.HealthCheckResponse{
		Status: res,
	}, nil
}

func (s *Server) RegisterUser(ctx context.Context, req *identityv1.RegisterUserRequest) (*identityv1.RegisterUserResponse, error) {
	res, err := s.identityService.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &identityv1.RegisterUserResponse{
		UserId: res.ID,
	}, nil
}
