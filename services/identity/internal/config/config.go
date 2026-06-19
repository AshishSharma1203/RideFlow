package config

import (
	"fmt"
	"time"
)

type Config struct {
	Server   ServerConfig
	Security SecurityConfig
	Postgres PostgresConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	GRPCPort int
}

type SecurityConfig struct {
	BcryptCost int
}

type JWTConfig struct {
	SecretKey              string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

func (c Config) Validate() error {
	if err := c.Server.Validate(); err != nil {
		return err
	}

	if err := c.Security.Validate(); err != nil {
		return err
	}

	if err := c.Postgres.Validate(); err != nil {
		return err
	}

	if err := c.JWT.Validate(); err != nil {
		return err
	}

	return nil
}

func (s ServerConfig) Validate() error {
	if s.GRPCPort < 1 || s.GRPCPort > 65535 {
		return fmt.Errorf("grpc port %d out of legal bounds", s.GRPCPort)
	}

	return nil
}

func (s SecurityConfig) Validate() error {
	if s.BcryptCost < 4 || s.BcryptCost > 31 {
		return fmt.Errorf("bcrypt cost %d out of legal security bounds", s.BcryptCost)
	}

	return nil
}

func (j JWTConfig) Validate() error {
	if j.SecretKey == "" {
		return fmt.Errorf("jwt secret key must not be empty")
	}
	if j.AccessTokenExpiration <= 0 {
		return fmt.Errorf("jwt access token expiration must be greater than 0")
	}
	if j.RefreshTokenExpiration <= 0 {
		return fmt.Errorf("jwt refresh token expiration must be greater than 0")
	}

	return nil
}
