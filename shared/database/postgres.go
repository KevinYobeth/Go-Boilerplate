package database

import (
	"database/sql"
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/log"

	_ "github.com/lib/pq"
)

type PostgresDB SqlDatabase

func InitPostgres() PostgresDB {
	log := log.InitLogger()
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
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(cfg.PostgresConnMaxLifeTime)
	db.SetConnMaxIdleTime(cfg.PostgresConnMaxIdleTime)
	db.SetMaxOpenConns(cfg.PostgresMaxOpenConns)
	db.SetMaxIdleConns(cfg.PostgresMaxIdleConns)

	err = db.Ping()
	if err != nil {
		log.Fatal(errors.NewInitializationError(err, "postgres").Message)
	}

	wrappedDB := NewDB(db)

	return wrappedDB
}
