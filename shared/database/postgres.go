package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/log"

	_ "github.com/lib/pq"
)

type PostgresDB SqlDatabase

func InitPostgres() PostgresDB {
	logger := log.InitLogger()
	cfg := config.LoadPostgresDBConfig()

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresUsername,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDBName,
		cfg.PostgresSSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal(err)
	}

	db.SetConnMaxLifetime(cfg.PostgresConnMaxLifeTime)
	db.SetConnMaxIdleTime(cfg.PostgresConnMaxIdleTime)
	db.SetMaxOpenConns(cfg.PostgresMaxOpenConns)
	db.SetMaxIdleConns(cfg.PostgresMaxIdleConns)

	err = db.Ping()
	if err != nil {
		logger.Fatal(errors.NewInitializationError(err, "postgres").Message)
	}

	wrappedDB := NewDB(db)

	wrappedDB.AddBeforeFunc(func(ctx context.Context, query string, args ...interface{}) context.Context {
		ctx = context.WithValue(ctx, constants.ContextKeyQueryStartKey, time.Now())
		return ctx
	})

	wrappedDB.AddAfterFunc(func(ctx context.Context, err error, query string, args ...interface{}) {
		query = strings.ReplaceAll(query, "\n", " ")
		query = strings.ReplaceAll(query, "\t", " ")

		logger := log.WithTrace(ctx, logger).With(
			"exec_time_elapsed", GetRequestDuration(ctx)/time.Millisecond,
			"query", query,
			"args", args,
			"caller", ctx.Value(constants.ContextKeySqlWrapperCaller),
		)

		logger.Debug("after executing query")

		if err != nil {
			logger.With("error", err).Error("error on executing query")
		}
	})

	return wrappedDB
}

func GetRequestDuration(ctx context.Context) time.Duration {
	now := time.Now()

	start := ctx.Value(constants.ContextKeyQueryStartKey)
	if start == nil {
		return 0
	}

	startTime, ok := start.(time.Time)
	if !ok {
		return 0
	}

	return now.Sub(startTime)
}
