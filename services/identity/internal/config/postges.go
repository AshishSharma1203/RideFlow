package config

import "fmt"

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (p PostgresConfig) Validate() error {
	if p.Host == "" {
		return fmt.Errorf("postgres host must not be empty")
	}

	if p.Port <= 0 {
		return fmt.Errorf("postgres port must be greater than 0")
	}

	if p.User == "" {
		return fmt.Errorf("postgres user must not be empty")
	}

	if p.Database == "" {
		return fmt.Errorf("postgres database must not be empty")
	}
	
	if p.Password == "" {
		return fmt.Errorf("postgres password must not be empty")
	}

	return nil
}