package application

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddress    string
	PostgresAddress string
	ServerPort      uint16
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress:    "localhost:6379",
		PostgresAddress: "postgres://postgres:postgres@localhost/library?sslmode=disable",
		ServerPort:      8080,
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if redisAddress, exists := os.LookupEnv("REDIS_ADDRESS"); exists {
		cfg.RedisAddress = redisAddress
	}

	if postgresAddress, exists := os.LookupEnv("POSTGRES_ADDRESS"); exists {
		cfg.PostgresAddress = postgresAddress
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
