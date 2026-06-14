package config

import "fmt"

type Config struct {
	Server   ServerConfig
	Security SecurityConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	GRPCPort int
}

type SecurityConfig struct {
	BcryptCost int
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