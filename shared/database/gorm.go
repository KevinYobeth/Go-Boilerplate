package database

import (
	"go-boilerplate/config"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormAbstraction struct {
	*gorm.Config
	Error        error
	RowsAffected int64
	Statement    *gorm.Statement
	clone        int
}

func InitGorm() *gorm.DB {
	appLog := log.InitLogger()
	gormLogger := log.NewGormLogger(100*time.Millisecond, logger.Info)
	cfg := config.LoadPostgresDBConfig()

	dsn := "host=localhost user=postgres password=postgres dbname=boilerplate port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		appLog.Fatal(err)
	}

	sql, err := db.DB()
	if err != nil {
		appLog.Fatal(err)
	}

	sql.SetConnMaxLifetime(cfg.PostgresConnMaxLifeTime)
	sql.SetConnMaxIdleTime(cfg.PostgresConnMaxIdleTime)
	sql.SetMaxOpenConns(cfg.PostgresMaxOpenConns)
	sql.SetMaxIdleConns(cfg.PostgresMaxIdleConns)

	db.Callback().Create().Before("gorm:create").Register("global_transaction_callback", GlobalTransactionCallback)
	db.Callback().Update().Before("gorm:update").Register("global_transaction_callback", GlobalTransactionCallback)
	db.Callback().Delete().Before("gorm:delete").Register("global_transaction_callback", GlobalTransactionCallback)

	err = sql.Ping()
	if err != nil {
		appLog.Fatal(errors.NewInitializationError(err, "gorm").Message)
	}

	return db
}
