package config

import "github.com/spf13/viper"

type DBConfig struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBString string `mapstructure:"DB_STRING"`
}

func LoadDBConfig() (config DBConfig, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
