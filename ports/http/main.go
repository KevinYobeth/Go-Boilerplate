package http

import (
	"context"
	"go-boilerplate/config"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/graceroutine"
	"go-boilerplate/shared/log"
	"go-boilerplate/shared/response"
	"go-boilerplate/shared/telemetry"
	"go-boilerplate/shared/types"
	"go-boilerplate/shared/utils"
	authorsHTTP "go-boilerplate/src/authors/presentation/http"
	authorIntraprocess "go-boilerplate/src/authors/presentation/intraprocess"
	authorsService "go-boilerplate/src/authors/services"
	booksAuthorIntraprocess "go-boilerplate/src/books/infrastructure/intraprocess"
	booksHTTP "go-boilerplate/src/books/presentation/http"
	booksService "go-boilerplate/src/books/services"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ztrue/tracerr"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func RunHTTPServer() {
	app := echo.New()

	appConfig := config.LoadAppConfig()
	logger := log.InitLogger()
	shutdownOtel, err := telemetry.InitOtel(context.Background())
	if err != nil {
		panic(err)
	}

	if strings.ToUpper(appConfig.AppEnv) == constants.APP_DEVELOPMENT {
		app.Debug = true
	}

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(otelecho.Middleware(appConfig.AppName))
	app.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, requestID string) {
			stdCtx := c.Request().Context()
			newCtx := context.WithValue(stdCtx, constants.ContextKeyRequestID, requestID)

			c.Set(string(constants.ContextKeyRequestID), requestID)
			c.SetRequest(c.Request().WithContext(newCtx))
		},
		Generator: func() string {
			return utils.RandomString(16)
		},
	}))
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
				fields = append(fields, "error", v.Error)
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
		return c.JSON(http.StatusOK, types.ResponseBody{
			Message: "ok",
		})
	})

	authorsService := authorsService.NewAuthorService()
	authorIntraprocess := authorIntraprocess.NewAuthorIntraprocessService(authorsService)

	booksAuthorIntraprocess := booksAuthorIntraprocess.NewBookAuthorIntraprocessService(authorIntraprocess)
	booksService := booksService.NewBookService(booksAuthorIntraprocess)
	booksServer := booksHTTP.NewBooksHTTPServer(&booksService)

	authorsServer := authorsHTTP.NewAuthorsHTTPServer(&authorsService)

	api := app.Group("/api")

	api.GET("/sandbox", func(c echo.Context) error {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return nil
	})

	booksServer.RegisterHTTPRoutes(api)
	authorsServer.RegisterHTTPRoutes(api)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		host := config.ServerHost + ":" + config.ServerHTTPPort
		if err := app.Start(host); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	<-signals

	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}

	shutdownOtel(context.Background())
	graceroutine.Stop()
	graceroutine.Wait()

	logger.Info("Server Shutdown")
}
