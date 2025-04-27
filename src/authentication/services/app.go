package services

import (
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/log"
	"go-boilerplate/src/authentication/infrastructure/repository"
	"go-boilerplate/src/authentication/services/command"
	"go-boilerplate/src/authentication/services/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Register command.RegisterHandler
}

type Queries struct {
	Login   query.LoginHandler
	GetUser query.GetUserHandler
}

func NewAuthenticationService() Application {
	db := database.InitPostgres()
	logger := log.InitLogger()

	repository := repository.NewAuthenticationPostgresRepository(db)

	return Application{
		Commands: Commands{
			Register: command.NewRegisterHandler(repository, logger),
		},
		Queries: Queries{
			Login:   query.NewLoginHandler(repository, logger),
			GetUser: query.NewGetUserHandler(repository, logger),
		},
	}
}
