package grpc

import (
	"context"
	"errors"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/dto"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	res, err := s.identityService.RegisterUser(ctx, dto.RegisterUserInput{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, mapRegisterUserError(err)
	}
	return &identityv1.RegisterUserResponse{
		UserId: res.UserID,
	}, nil
}

func mapRegisterUserError(err error) error {
	switch {
	case errors.Is(err, service.ErrInvalidUsername),
		errors.Is(err, service.ErrInvalidEmail),
		errors.Is(err, service.ErrInvalidPassword):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, service.ErrEmailAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
