package config

type ServerConfig struct {
	ServerHost string `env:"SERVER_HOST" default:"localhost"`
	ServerType string `env:"SERVER_TYPE" default:"http"`

	ServerHTTPPort string `env:"SERVER_HTTP_PORT" default:"8080"`
	ServerGRPCPort string `env:"SERVER_GRPC_PORT" default:"8181"`
}

var serverConfig ServerConfig

func LoadServerConfig() ServerConfig {
	if serverConfig != (ServerConfig{}) {
		return serverConfig
	}

	loadConfig(&serverConfig)

	return serverConfig
}
