package config

import (
	"log"
)

type Config struct {
	Database DBConfig
	Server   ServerConfig
}

func InitConfig() (config Config) {
	dbConfig, err := LoadDBConfig()
	if err != nil {
		log.Fatal("cannot load database config:", err)
	}

	serverConfig, err := LoadServerConfig()
	if err != nil {
		log.Fatal("cannot load server config:", err)
	}

	return Config{
		Database: dbConfig,
		Server:   serverConfig,
	}
}
