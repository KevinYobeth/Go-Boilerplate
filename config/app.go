package config

type AppConfig struct {
	AppEnv string `env:"APP_ENV" default:"development"`
}

var appConfig AppConfig

func LoadAppConfig() AppConfig {
	if appConfig != (AppConfig{}) {
		return appConfig
	}

	loadConfig(&appConfig)

	return appConfig
}
