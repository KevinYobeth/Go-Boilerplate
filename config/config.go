package config

import (
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Database PostgresConfig
	Server   ServerConfig
}

func validateConfig(dest interface{}) error {
	validate := validator.New()

	err := validate.Struct(dest)
	if err != nil {
		return err
	}

	return nil
}

func loadConfig(dest interface{}) error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&dest, func(config *mapstructure.DecoderConfig) {
		config.TagName = "env"
	})
	if err != nil {
		panic(err)
	}

	err = validateConfig(dest)
	if err != nil {
		panic(err)
	}

	return nil
}

func InitConfig() Config {
	dbConfig := LoadPostgresDBConfig()
	serverConfig := LoadServerConfig()

	return Config{
		Database: dbConfig,
		Server:   serverConfig,
	}
}
