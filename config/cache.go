package config

type RedisConfig struct {
	RedisHost     string `env:"REDIS_HOST" default:"localhost" validate:"required"`
	RedisPort     string `env:"REDIS_PORT" default:"6379" validate:"required"`
	RedisPassword string `env:"REDIS_PASS" default:""`
	RedisDB       int    `env:"REDIS_DB" default:"0"`
}

var redisConfig RedisConfig

func LoadRedisCacheConfig() RedisConfig {
	if redisConfig != (RedisConfig{}) {
		return redisConfig
	}

	loadConfig(&redisConfig)

	return redisConfig
}
