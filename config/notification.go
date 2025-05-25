package config

type NotificationConfig struct {
	NotificationStrategy string `env:"NOTIFICATION_STRATEGY" default:"email" validate:"required,oneof=email database"`
}

var notificationConfig NotificationConfig

func LoadNotificationConfig() NotificationConfig {
	if notificationConfig != (NotificationConfig{}) {
		return notificationConfig
	}

	loadConfig(&notificationConfig)

	return notificationConfig
}
