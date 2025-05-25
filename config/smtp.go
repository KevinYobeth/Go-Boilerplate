package config

type SMTPConfig struct {
	SMTPHost     string `env:"SMTP_HOST" default:"localhost" validate:"required"`
	SMTPPort     int    `env:"SMTP_PORT" default:"1025" validate:"required"`
	SMTPUsername string `env:"SMTP_USERNAME" default:"mailpit" validate:"required"`
	SMTPPassword string `env:"SMTP_PASSWORD" default:"mailpit" validate:"required"`
}

var smtpConfig SMTPConfig

func LoadSMTPConfig() SMTPConfig {
	if smtpConfig != (SMTPConfig{}) {
		return smtpConfig
	}

	loadConfig(&smtpConfig)

	return smtpConfig
}
