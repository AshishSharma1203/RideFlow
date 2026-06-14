package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	_ = godotenv.Load()

	grpcPortStr := os.Getenv(EnvGRPCPort)
	bcryptCostStr := os.Getenv(EnvBcryptCost)

	postgresHost := os.Getenv(EnvPostgresHost)
	postgresPortStr := os.Getenv(EnvPostgresPort)
	postgresUser := os.Getenv(EnvPostgresUser)
	postgresPassword := os.Getenv(EnvPostgresPassword)
	postgresDatabase := os.Getenv(EnvPostgresDatabase)

	grpcPort, err := strconv.Atoi(grpcPortStr)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid %s '%s': %w",
			EnvGRPCPort,
			grpcPortStr,
			err,
		)
	}

	bcryptCost, err := strconv.Atoi(bcryptCostStr)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid %s '%s': %w",
			EnvBcryptCost,
			bcryptCostStr,
			err,
		)
	}

	postgresPort, err := strconv.Atoi(postgresPortStr)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid %s '%s': %w",
			EnvPostgresPort,
			postgresPortStr,
			err,
		)
	}

	cfg := &Config{
		Server: ServerConfig{
			GRPCPort: grpcPort,
		},
		Security: SecurityConfig{
			BcryptCost: bcryptCost,
		},
		Postgres: PostgresConfig{
			Host:     postgresHost,
			Port:     postgresPort,
			User:     postgresUser,
			Password: postgresPassword,
			Database: postgresDatabase,
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}
