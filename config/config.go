package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database
}

func LoadDBConfig() (db Database, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&db)
	return
}

func InitConfig() (config Config) {
	dbConfig, err := LoadDBConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return Config{
		Database: dbConfig,
	}
}
