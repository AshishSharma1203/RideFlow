package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/model"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/repository/sqlc"
)

var ErrUserNotFound = errors.New("user not found")

type PostgresUserRepository struct {
	queries *sqlc.Queries
}

func NewPostgresUserRepository(
	queries *sqlc.Queries,
) *PostgresUserRepository {

	return &PostgresUserRepository{
		queries: queries,
	}
}

func (r *PostgresUserRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*model.User, error) {

	user, err := r.queries.GetUserByEmail(
		ctx,
		email,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return mapSqlUserToModelUser(user), nil
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	// params := sqlc.CreateUserParams{
	// 	Name:         user.Name,
	// 	Email:        user.Email,
	// 	PasswordHash: user.Password,
	// }
	// return r.queries.CreateUser(ctx, params)
	return nil
}

func mapSqlUserToModelUser(
	sqlUser sqlc.User,
) *model.User {

	return &model.User{
		ID:           sqlUser.ID.String(),
		Name:         sqlUser.Name,
		Email:        sqlUser.Email,
		PasswordHash: sqlUser.PasswordHash,
	}
}
