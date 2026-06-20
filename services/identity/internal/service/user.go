package service

import (
	"context"
	"errors"
	"github.com/ashishSharma1203/rideflow/pkg/auth"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/dto"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/model"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/repository"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/security"
)

type IdentityService struct {
	userRepo     repository.UserRepository
	hasher       security.PasswordHasher
	tokenManager auth.TokenManager
}

func NewIdentityService(userRepo repository.UserRepository, hasher security.PasswordHasher, tokenManager auth.TokenManager) *IdentityService {
	return &IdentityService{
		userRepo:     userRepo,
		hasher:       hasher,
		tokenManager: tokenManager,
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

func (s *IdentityService) LoginUser(ctx context.Context, user dto.LoginUserInput) (*dto.LoginUserOutput, error) {
	if !isValidEmailFormat(user.Email) {
		return nil, ErrInvalidEmail
	}
	userModel, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if err := s.hasher.Compare(userModel.PasswordHash, user.Password); err != nil {
		return nil, ErrInvalidPassword
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(userModel.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenManager.GenerateRefreshToken(userModel.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginUserOutput{
		ID:           userModel.ID,
		Username:     userModel.Username,
		Email:        userModel.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
