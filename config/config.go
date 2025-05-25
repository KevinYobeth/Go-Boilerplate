package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	OpenTelemetry OpenTelemetryConfig
	Postgres      PostgresConfig
	Redis         RedisConfig
	Cache         CacheConfig
	Server        ServerConfig
	App           AppConfig
	Event         RabbitMQConfig
	Setting       SettingConfig
	Notification  NotificationConfig
	SMTP          SMTPConfig
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
	return Config{
		Postgres:      LoadPostgresDBConfig(),
		Redis:         LoadRedisCacheConfig(),
		Server:        LoadServerConfig(),
		App:           LoadAppConfig(),
		Event:         LoadRabbitMQConfig(),
		OpenTelemetry: LoadOpenTelemetryConfig(),
		Setting:       LoadSettingConfig(),
		Cache:         LoadCacheConfig(),
		Notification:  LoadNotificationConfig(),
		SMTP:          LoadSMTPConfig(),
	}
}
