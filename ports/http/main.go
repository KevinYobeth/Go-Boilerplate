package http

import (
	"go-boilerplate/config"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/log"
	books "go-boilerplate/src/books/infrastructure/transport"
	"go-boilerplate/src/books/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ztrue/tracerr"
)

func RunHTTPServer() {
	app := echo.New()

	appConfig := config.LoadAppConfig()
	logger := log.InitLogger()

	if appConfig.AppEnv == constants.APP_DEVELOPMENT {
		app.Debug = true
	}

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogMethod:    true,
		LogRequestID: true,
		LogHost:      true,
		LogRemoteIP:  true,
		LogError:     true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {

			fields := []interface{}{
				"URI", v.URI,
				"status", v.Status,
				"method", v.Method,
				"latency", v.Latency,
				"host", v.Host,
				"remote_ip", v.RemoteIP,
				"request_id", v.RequestID,
			}

			if v.Error != nil {
				fields = append(fields, "trace", tracerr.StackTrace(errors.GetTracerrErr(v.Error)))
			}

			s := v.Status
			switch {
			case s >= 500:
				logger.Errorw("request", fields...)
			case s >= 400:
				logger.Warnw("request", fields...)
			default:
				logger.Infow("request", fields...)
			}

			return nil
		},
	}))

	config := config.LoadServerConfig()

	app.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	booksService := services.NewBookService()
	booksServer := books.NewHTTPServer(&booksService)

	api := app.Group("/api")

	booksServer.RegisterBookHTTPRoutes(api)

	logger.Fatal(app.Start(config.ServerHost + ":" + config.ServerPort))
}
