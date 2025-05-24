package config

import (
	"time"
)

type PostgresConfig struct {
	PostgresUsername        string        `env:"POSTGRES_USERNAME" default:"postgres" validate:"required"`
	PostgresPassword        string        `env:"POSTGRES_PASSWORD" default:"postgres" validate:"required"`
	PostgresHost            string        `env:"POSTGRES_HOST" default:"localhost" validate:"required"`
	PostgresPort            string        `env:"POSTGRES_PORT" default:"5432" validate:"required"`
	PostgresDBName          string        `env:"POSTGRES_DB_NAME" default:"default" validate:"required"`
	PostgresSSLMode         string        `env:"POSTGRES_SSL_MODE" default:"false" validate:"required"`
	PostgresMaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS" default:"30" validate:"required"`
	PostgresMaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS" default:"6" validate:"required"`
	PostgresConnMaxLifeTime time.Duration `env:"POSTGRES_CONN_MAX_LIFE_TIME" default:"30m" validate:"required"`
	PostgresConnMaxIdleTime time.Duration `env:"POSTGRES_CONN_MAX_IDLE_TIME" default:"0"`
}

var postgresConfig PostgresConfig

func LoadPostgresDBConfig() PostgresConfig {
	if postgresConfig != (PostgresConfig{}) {
		return postgresConfig
	}

	loadConfig(&postgresConfig)

	return postgresConfig
}
