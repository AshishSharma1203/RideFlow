package repository

import (
	"context"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/model"
)

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		user *model.User,
	) (*model.User, error)

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (*model.User, error)
}
