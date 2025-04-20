package services

import (
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/log"
	"go-boilerplate/src/authentication/infrastructure/repository"
	"go-boilerplate/src/authentication/services/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Register command.RegisterHandler
}

type Queries struct {
}

func NewAuthenticationService() Application {
	db := database.InitPostgres()
	logger := log.InitLogger()

	repository := repository.NewAuthenticationPostgresRepository(db)

	return Application{
		Commands: Commands{
			Register: command.NewRegisterHandler(repository, logger),
		},
		Queries: Queries{},
	}
}
