package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort   int
	BcryptCost int
}

func Load() (*Config, error) {
	// 1. load the env files
	_ = godotenv.Load()

	// 2. Read the raw string values from the OS environment
	grpcPortStr := os.Getenv("IDENTITY_GRPC_PORT")
	bcryptCostStr := os.Getenv("IDENTITY_BCRYPT_COST")

	// 3. Convert GRPCPort string to int
	grpcPort, err := strconv.Atoi(grpcPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid IDENTITY_GRPC_PORT '%s': %w", grpcPortStr, err)
	}

	// 4. Convert BcryptCost string to int
	bcryptCost, err := strconv.Atoi(bcryptCostStr)
	if err != nil {
		return nil, fmt.Errorf("invalid IDENTITY_BCRYPT_COST '%s': %w", bcryptCostStr, err)
	}

	// 3. Populate the values
	cfg := &Config{
		GRPCPort:   grpcPort,
		BcryptCost: bcryptCost,
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}
	return cfg, nil

}

func (cfg *Config) Validate() error {

	// check errors
	if cfg.GRPCPort < 1 || cfg.GRPCPort > 65335 {
		return fmt.Errorf("grpc port %d out of legal bound", cfg.GRPCPort)
	}
	if cfg.BcryptCost < 4 || cfg.BcryptCost > 31 {
		return fmt.Errorf("bcrypt cost %d out of legal security bounds ", cfg.BcryptCost)
	}

	return nil
}
