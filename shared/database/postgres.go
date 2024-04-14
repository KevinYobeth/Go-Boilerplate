package database

import (
	"database/sql"
	"fmt"
	"go-boilerplate/config"
	"log"

	_ "github.com/lib/pq"
)

func InitPostgres() *sql.DB {
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

	return db
}
