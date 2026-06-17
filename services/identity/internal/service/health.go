package service

import (
	"context"
	"errors"
	"strings"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/model"
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

func (s *IdentityService) RegisterUser(ctx context.Context, req *identityv1.RegisterUserRequest) (*model.User, error) {
	username := strings.TrimSpace(req.GetUsername())
	email := strings.TrimSpace(req.GetEmail())
	password := req.GetPassword()

	if username == "" {
		return nil, errors.New("username is required")
	}

	if !isValidEmail(email) {
		return nil, errors.New("invalid email")
	}

	if password == "" {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	userModel := model.User{
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	return s.userRepo.CreateUser(ctx, &userModel)
}

func isValidEmail(email string) bool {
	return email != ""
}
