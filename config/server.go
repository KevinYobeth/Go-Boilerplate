package config

type ServerConfig struct {
	ServerPort string `env:"SERVER_PORT" default:"8080"`
	ServerType string `env:"SERVER_TYPE" default:"http"`
}

var serverConfig ServerConfig

func LoadServerConfig() ServerConfig {
	if serverConfig != (ServerConfig{}) {
		return serverConfig
	}

	loadConfig(&serverConfig)

	return serverConfig
}
