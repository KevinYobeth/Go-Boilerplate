package http

import (
	"context"
	"go-boilerplate/config"
	"go-boilerplate/shared/cache"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/event"
	"go-boilerplate/shared/graceroutine"
	"go-boilerplate/shared/log"
	"go-boilerplate/shared/types"
	"go-boilerplate/src/authors/infrastructure/repository"
	authorsHTTP "go-boilerplate/src/authors/presentation/http"
	"go-boilerplate/src/authors/presentation/intraprocess"
	authorsService "go-boilerplate/src/authors/services"
	booksRepository "go-boilerplate/src/books/infrastructure/repository"
	booksHTTP "go-boilerplate/src/books/presentation/http"
	booksService "go-boilerplate/src/books/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	lru := cache.InitRedis()
	db := database.InitPostgres()
	manager := database.NewTransactionManager(db)

	bookRepo := booksRepository.NewBooksPostgresRepository(db)
	bookCache := booksRepository.NewBooksRedisCache(lru)
	// authorRepo := authorsRepository.NewAuthorsPostgresRepository(db)

	authorRepo := repository.NewAuthorsPostgresRepository(db)
	authorPublisher := event.InitPublisher(event.PublisherOptions{
		Topic: "authors",
	})

	authorsService := authorsService.NewAuthorService(authorRepo, authorPublisher)
	authorIntraprocess := intraprocess.NewAuthorIntraprocessService(authorsService)

	booksService := booksService.NewBookService(bookRepo, bookCache, manager, authorIntraprocess)
	booksServer := booksHTTP.NewBooksHTTPServer(&booksService)

	authorsServer := authorsHTTP.NewAuthorsHTTPServer(&authorsService)

	api := app.Group("/api")

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

	graceroutine.Stop()
	graceroutine.Wait()

	logger.Info("Server Shutdown")
}
