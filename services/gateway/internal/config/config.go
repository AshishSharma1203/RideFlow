package config

import "fmt"

type Config struct {
	Server   ServerConfig
	Identity IdentityConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	HTTPPort int
}

type IdentityConfig struct {
	GRPCAddr string
}

type JWTConfig struct {
	SecretKey string
}

func (c Config) Validate() error {
	if err := c.Server.Validate(); err != nil {
		return err
	}

	if err := c.Identity.Validate(); err != nil {
		return err
	}

	if err := c.JWT.Validate(); err != nil {
		return err
	}

	return nil
}

func (s ServerConfig) Validate() error {
	if s.HTTPPort < 1 || s.HTTPPort > 65535 {
		return fmt.Errorf("http port %d out of legal bounds", s.HTTPPort)
	}

	return nil
}

func (i IdentityConfig) Validate() error {
	if i.GRPCAddr == "" {
		return fmt.Errorf("identity grpc address must not be empty")
	}

	return nil
}

func (j JWTConfig) Validate() error {
	if j.SecretKey == "" {
		return fmt.Errorf("jwt secret key must not be empty")
	}

	return nil
}
