package service

import (
	"context"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/client"
)

type RegisterUserInput struct {
	Username string
	Email    string
	Password string
}

type RegisterUserOutput struct {
	UserID string
}

type IdentityService struct {
	identityClient *client.IdentityClient
}

func NewIdentityService(identityClient *client.IdentityClient) *IdentityService {
	return &IdentityService{
		identityClient: identityClient,
	}
}

func (s *IdentityService) RegisterUser(
	ctx context.Context,
	input RegisterUserInput,
) (*RegisterUserOutput, error) {
	resp, err := s.identityClient.Client.RegisterUser(
		ctx,
		&identityv1.RegisterUserRequest{
			Username: input.Username,
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	return &RegisterUserOutput{
		UserID: resp.GetUserId(),
	}, nil
}
