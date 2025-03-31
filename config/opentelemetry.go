package config

import "time"

type OpenTelemetryConfig struct {
	OtelDisabled             bool          `env:"OTEL_DISABLED" default:"false"`
	OtelGRPCEndpoint         string        `env:"OTEL_GRPC_ENDPOINT" default:"localhost:4317" validate:"required"`
	OtelRetryEnabled         bool          `env:"OTEL_RETRY_ENABLED" default:"true" validate:"required"`
	OtelRetryInitialInterval time.Duration `env:"OTEL_RETRY_INITIAL_INTERVAL" default:"5s" validate:"required"`
	OtelRetryMaxInterval     time.Duration `env:"OTEL_RETRY_MAX_INTERVAL" default:"60s" validate:"required"`
	OtelRetryMaxElapsedTime  time.Duration `env:"OTEL_RETRY_MAX_ELAPSED_TIME" default:"300s" validate:"required"`
}

var openTelemetryConfig OpenTelemetryConfig

func LoadOpenTelemetryConfig() OpenTelemetryConfig {
	if openTelemetryConfig != (OpenTelemetryConfig{}) {
		return openTelemetryConfig
	}

	loadConfig(&openTelemetryConfig)

	return openTelemetryConfig
}
