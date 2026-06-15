package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	pool, err := pgxpool.New(
		context.Background(),
		dsn,
	)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return pool, nil
}
