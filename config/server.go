package config

import "github.com/spf13/viper"

type ServerConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerType string `mapstructure:"SERVER_TYPE"`
}

func LoadServerConfig() (config ServerConfig, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
