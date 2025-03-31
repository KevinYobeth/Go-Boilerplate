package cache

import (
	"context"
	"go-boilerplate/config"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/telemetry"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	ctx := context.Background()

	redisConfig := config.LoadRedisCacheConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.RedisHost + ":" + redisConfig.RedisPort,
		Password: redisConfig.RedisPassword,
		DB:       redisConfig.RedisDB,
	})

	client.AddHook(telemetry.NewRedisTracingHook())

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(errors.NewInitializationError(err, "redis").Message)
	}

	return client
}
