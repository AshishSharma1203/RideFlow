package config

const (
	// Server
	EnvGRPCPort = "IDENTITY_GRPC_PORT"

	// Security
	EnvBcryptCost = "IDENTITY_BCRYPT_COST"

	// Database
	EnvPostgresHost     = "POSTGRES_HOST"
	EnvPostgresPort     = "POSTGRES_PORT"
	EnvPostgresUser     = "POSTGRES_USER"
	EnvPostgresPassword = "POSTGRES_PASSWORD"
	EnvPostgresDatabase = "POSTGRES_DATABASE"
)
