package config

const (
	// Server
	EnvGRPCPort = "IDENTITY_GRPC_PORT"

	// Security
	EnvBcryptCost = "IDENTITY_BCRYPT_COST"

	// Database
	EnvPostgresHost     = "IDENTITY_POSTGRES_HOST"
	EnvPostgresPort     = "IDENTITY_POSTGRES_PORT"
	EnvPostgresUser     = "IDENTITY_POSTGRES_USER"
	EnvPostgresPassword = "IDENTITY_POSTGRES_PASSWORD"
	EnvPostgresDatabase = "IDENTITY_POSTGRES_DATABASE"
)