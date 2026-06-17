package service

import (
	"context"
	"errors"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/dto"
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

func (s *IdentityService) RegisterUser(
	ctx context.Context,
	input dto.RegisterUserInput,
) (*dto.RegisterUserOutput, error) {
	input = normalizeRegisterUserInput(input)

	if err := validateRegisterUserInput(input); err != nil {
		return nil, err
	}

	existingUser, err := s.userRepo.GetUserByEmail(ctx, input.Email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return nil, err
	}

	hashedPassword, err := s.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	userModel := model.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: hashedPassword,
	}

	createdUser, err := s.userRepo.CreateUser(ctx, &userModel)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserOutput{
		UserID: createdUser.ID,
	}, nil
}
