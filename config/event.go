package config

type RabbitMQConfig struct {
	RabbitMQUsername string `env:"RABBITMQ_USERNAME" default:"rabbitmq" validate:"required"`
	RabbitMQPassword string `env:"RABBITMQ_PASSWORD" default:"rabbitmq" validate:"required"`
	RabbitMQHost     string `env:"RABBITMQ_HOST" default:"localhost" validate:"required"`
	RabbitMQPort     string `env:"RABBITMQ_PORT" default:"5672" validate:"required"`
}

var rabbitMQConfig RabbitMQConfig

func LoadRabbitMQConfig() RabbitMQConfig {
	if rabbitMQConfig != (RabbitMQConfig{}) {
		return rabbitMQConfig
	}

	loadConfig(&rabbitMQConfig)

	return rabbitMQConfig
}
