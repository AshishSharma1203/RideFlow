package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	_ = godotenv.Load()

	httpPortStr := os.Getenv(EnvHTTPPort)
	identityGRPCAddr := os.Getenv(EnvIdentityGRPCAddr)
	jwtSecretKey := os.Getenv(EnvJWTSecretKey)

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid %s '%s': %w",
			EnvHTTPPort,
			httpPortStr,
			err,
		)
	}

	cfg := &Config{
		Server: ServerConfig{
			HTTPPort: httpPort,
		},
		Identity: IdentityConfig{
			GRPCAddr: identityGRPCAddr,
		},
		JWT: JWTConfig{
			SecretKey: jwtSecretKey,
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}
