package service

import (
	"context"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/security"
)

type IdentityService struct {
	hasher security.PasswordHasher
}

func NewIdentityService(hasher security.PasswordHasher) *IdentityService {
	return &IdentityService{
		hasher: hasher,
	}
}

func (s *IdentityService) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}
