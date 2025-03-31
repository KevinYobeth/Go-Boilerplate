package config

type AppConfig struct {
	AppEnv     string `env:"APP_ENV" default:"development"`
	AppName    string `env:"APP_NAME" default:"boilerplate"`
	AppVersion string `env:"APP_VERSION" default:"0.0.1"`
}

var appConfig AppConfig

func LoadAppConfig() AppConfig {
	if appConfig != (AppConfig{}) {
		return appConfig
	}

	loadConfig(&appConfig)

	return appConfig
}
