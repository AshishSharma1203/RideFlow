package service

import (
	"context"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/repository"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/security"
)

type IdentityService struct {
	userRepo repository.UserRepository
	hasher   security.PasswordHasher
}

func NewIdentityService(userRepo repository.UserRepository, hasher security.PasswordHasher) *IdentityService {
	return &IdentityService{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (s *IdentityService) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}
