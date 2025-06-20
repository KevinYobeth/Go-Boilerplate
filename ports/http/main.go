package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/kevinyobeth/go-boilerplate/config"
	authenticationHTTP "github.com/kevinyobeth/go-boilerplate/internal/authentication/presentation/http"
	authenticationService "github.com/kevinyobeth/go-boilerplate/internal/authentication/services"
	linkHTTP "github.com/kevinyobeth/go-boilerplate/internal/link/presentation/http"
	linkService "github.com/kevinyobeth/go-boilerplate/internal/link/services"
	userHTTP "github.com/kevinyobeth/go-boilerplate/internal/user/presentation/http"
	userIntraprocess "github.com/kevinyobeth/go-boilerplate/internal/user/presentation/intraprocess"
	userService "github.com/kevinyobeth/go-boilerplate/internal/user/services"
	"github.com/kevinyobeth/go-boilerplate/shared/constants"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/graceroutine"
	"github.com/kevinyobeth/go-boilerplate/shared/log"
	"github.com/kevinyobeth/go-boilerplate/shared/response"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"
	"github.com/kevinyobeth/go-boilerplate/shared/types"
	"github.com/kevinyobeth/go-boilerplate/shared/utils"

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
				fields = append(fields, "error_raw", errors.GetGenericError(v.Error).Unwrap())
				fields = append(fields, "error_type", errors.GetGenericError(v.Error).Type)
				fields = append(fields, "error_metadata", errors.GetGenericError(v.Error).Metadata)
				fields = append(fields, "trace", tracerr.StackTrace(errors.GetTracerrErr(v.Error)))
			}

			if strings.ToUpper(appConfig.AppEnv) == constants.APP_DEVELOPMENT {
				tracerr.PrintSourceColor(errors.GetTracerrErr(v.Error), 3, 3)
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

	userService := userService.NewUserService()
	userServer := userHTTP.NewUserHTTPServer(&userService)
	userIntraprocess := userIntraprocess.NewUserIntraprocessService(userService)

	authenticationService := authenticationService.NewAuthenticationService(userIntraprocess)
	authenticationServer := authenticationHTTP.NewAuthenticationHTTPServer(&authenticationService)

	linkService := linkService.NewLinkService()
	linkServer := linkHTTP.NewLinkHTTPServer(&linkService)

	api := app.Group("/api")
	{
		authenticationServer.RegisterHTTPRoutes(api)
		userServer.RegisterHTTPRoutes(api)
		linkServer.RegisterHTTPRoutes(api, app)
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
