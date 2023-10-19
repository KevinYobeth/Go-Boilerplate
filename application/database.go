package application

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToPostgres() {
	conn, err := gorm.Open(postgres.Open(LoadConfig().PostgresAddress), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect to postgres :%w", err)
	}

	DB = conn
}

func GetPostgresDB() *gorm.DB {
	if DB == nil {
		connectToPostgres()
	}

	return DB
}
