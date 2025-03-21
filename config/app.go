package config

type AppConfig struct {
	AppEnv  string `env:"APP_ENV" default:"development"`
	AppName string `env:"APP_NAME" default:"boilerplate"`
}

var appConfig AppConfig

func LoadAppConfig() AppConfig {
	if appConfig != (AppConfig{}) {
		return appConfig
	}

	loadConfig(&appConfig)

	return appConfig
}
