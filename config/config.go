package config

import (
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Database PostgresConfig
	Server   ServerConfig
	App      AppConfig
}

func validateConfig(dest interface{}) error {
	validate := validator.New()

	err := validate.Struct(dest)
	if err != nil {
		return err
	}

	return nil
}

func loadConfig(dest interface{}) {
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
}

func InitConfig() Config {
	dbConfig := LoadPostgresDBConfig()
	serverConfig := LoadServerConfig()
	appConfig := LoadAppConfig()

	return Config{
		Database: dbConfig,
		Server:   serverConfig,
		App:      appConfig,
	}
}
