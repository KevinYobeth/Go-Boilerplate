package log

import (
	"go-boilerplate/config"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/errors"
	"log"
	"os"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	if Logger != nil {
		return Logger
	}

	appConfig := config.LoadAppConfig()

	if appConfig.AppEnv == constants.APP_DEVELOPMENT {
		logDir := "logs"
		cfg := zap.NewDevelopmentConfig()

		if err := os.MkdirAll(logDir, 0755); err != nil {
			log.Fatalf("failed to create log directory: %v", err)
		}

		cfg.OutputPaths = []string{"logs/development.log", "stderr"}

		newLogger, err := cfg.Build()

		if err != nil {
			log.Fatal(errors.NewInitializationError(err, "logger").Message)
			return nil
		}

		sugar := newLogger.Sugar()
		Logger = sugar

		sugar.Infow("logger initialized in development mode")

		return sugar
	}

	newLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(errors.NewInitializationError(err, "logger").Message)
		return nil
	}

	sugar := newLogger.Sugar()
	Logger = sugar

	sugar.Infow("logger initialized in production mode")

	return sugar
}
