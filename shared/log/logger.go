package log

import (
	"go-boilerplate/config"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/errors"
	"log"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	if Logger != nil {
		return Logger
	}

	appConfig := config.LoadAppConfig()

	if appConfig.AppEnv == constants.APP_DEVELOPMENT {
		newLogger, err := zap.NewDevelopment()

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
