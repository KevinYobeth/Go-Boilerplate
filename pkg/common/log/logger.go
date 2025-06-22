package log

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/constants"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/telemetry"

	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/log/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func WithTrace(ctx context.Context, logger *zap.SugaredLogger) *zap.SugaredLogger {
	fields := []any{}

	if traceID := telemetry.GetTraceID(ctx); traceID != "" {
		fields = append(fields, "trace_id", traceID)
	}
	if spanID := telemetry.GetSpanID(ctx); spanID != "" {
		fields = append(fields, "span_id", spanID)
	}

	return logger.With(fields...)
}

func InitLogger() *zap.SugaredLogger {
	if Logger != nil {
		return Logger
	}

	appConfig := config.LoadAppConfig()
	baseLogger, environment := setupBaseLogger(appConfig.AppEnv)

	core := baseLogger.Core()

	cfg := config.LoadOpenTelemetryConfig()
	if !cfg.OtelDisabled {
		loggerProvider := global.GetLoggerProvider()
		otelZapCore := otelzap.NewCore("github.com/kevinyobeth/go-boilerplate",
			otelzap.WithLoggerProvider(loggerProvider),
		)

		core = zapcore.NewTee(
			baseLogger.Core(),
			otelZapCore,
		)
	}
	newLogger := zap.New(core)

	newLogger.Info("logger initialized in " + environment)

	Logger = newLogger.Sugar()
	return Logger
}

func setupBaseLogger(appEnv string) (*zap.Logger, string) {
	if strings.ToUpper(appEnv) == constants.APP_DEVELOPMENT {
		return setupDevelopmentLogger()
	}
	return setupProductionLogger()
}

func setupDevelopmentLogger() (*zap.Logger, string) {
	logDir := "logs"
	cfg := zap.NewDevelopmentConfig()

	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("failed to create log directory: %v", err)
	}

	cfg.OutputPaths = []string{"logs/development.log", "stderr"}

	newLogger, err := cfg.Build()
	if err != nil {
		log.Fatal(errors.NewInitializationError(err, "logger").Message)
	}

	return newLogger, "development"
}

func setupProductionLogger() (*zap.Logger, string) {
	prodLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(errors.NewInitializationError(err, "logger").Message)
	}

	return prodLogger, "production"
}
