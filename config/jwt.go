package config

import "time"

type JWTConfig struct {
	JWTSecret    string        `env:"JWT_SECRET" validate:"required"`
	JWTShortLife time.Duration `env:"JWT_SHORT_LIFE" default:"15m" validate:"required"`
	JWTLongLife  time.Duration `env:"JWT_LONG_LIFE" default:"24h" validate:"required"`
}

var jwtConfig JWTConfig

func LoadJWTConfig() JWTConfig {
	if jwtConfig != (JWTConfig{}) {
		return jwtConfig
	}

	loadConfig(&jwtConfig)

	return jwtConfig
}
