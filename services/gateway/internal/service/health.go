package service

import (
	"context"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/client"
)

type HealthService struct {
	identityClient *client.IdentityClient
}

func NewHealthService(identityClient *client.IdentityClient) *HealthService {
	return &HealthService{
		identityClient: identityClient,
	}
}

func (s *HealthService) Check(ctx context.Context) (string, error) {
	resp, err := s.identityClient.Client.HealthCheck(
		ctx,
		&identityv1.HealthCheckRequest{},
	)
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}