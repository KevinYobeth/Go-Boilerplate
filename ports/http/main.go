package http

import (
	"context"
	"go-boilerplate/config"
	authenticationHTTP "go-boilerplate/internal/authentication/presentation/http"
	authenticationService "go-boilerplate/internal/authentication/services"
	authorsHTTP "go-boilerplate/internal/authors/presentation/http"
	authorIntraprocess "go-boilerplate/internal/authors/presentation/intraprocess"
	authorsService "go-boilerplate/internal/authors/services"
	booksAuthorIntraprocess "go-boilerplate/internal/books/infrastructure/intraprocess"
	booksHTTP "go-boilerplate/internal/books/presentation/http"
	booksService "go-boilerplate/internal/books/services"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/graceroutine"
	"go-boilerplate/shared/log"
	"go-boilerplate/shared/response"
	"go-boilerplate/shared/telemetry"
	"go-boilerplate/shared/types"
	"go-boilerplate/shared/utils"
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
	shutdownOtel, err := telemetry.InitOtel(context.Background())
	logger := log.InitLogger()

	if err != nil {
		panic(err)
	}

	if strings.ToUpper(appConfig.AppEnv) == constants.APP_DEVELOPMENT {
		app.Debug = true
	}

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(otelecho.Middleware(appConfig.AppName,
		otelecho.WithSkipper(func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "/health")
		})))
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
				"trace_id", telemetry.GetTraceID(c.Request().Context()),
				"span_id", telemetry.GetSpanID(c.Request().Context()),
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

	app.GET("/health", func(c echo.Context) error {
		response.SendHTTP(c, &types.Response{
			Body: types.ResponseBody{
				Data: map[string]any{
					"name":        appConfig.AppName,
					"version":     appConfig.AppVersion,
					"environment": appConfig.AppEnv,
				},
				Message: "ok",
			},
		})
		return nil
	})

	authenticationService := authenticationService.NewAuthenticationService()
	authenticationServer := authenticationHTTP.NewAuthenticationHTTPServer(&authenticationService)

	authorsService := authorsService.NewAuthorService()
	authorIntraprocess := authorIntraprocess.NewAuthorIntraprocessService(authorsService)

	booksAuthorIntraprocess := booksAuthorIntraprocess.NewBookAuthorIntraprocessService(authorIntraprocess)
	booksService := booksService.NewBookService(booksAuthorIntraprocess)
	booksServer := booksHTTP.NewBooksHTTPServer(&booksService)

	authorsServer := authorsHTTP.NewAuthorsHTTPServer(&authorsService)

	api := app.Group("/api")
	{
		authenticationServer.RegisterHTTPRoutes(api)
		booksServer.RegisterHTTPRoutes(api)
		authorsServer.RegisterHTTPRoutes(api)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	config := config.LoadServerConfig()

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
