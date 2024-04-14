package services

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
}

func NewAuthorService() Application {
	return Application{
		Commands: Commands{},
		Queries:  Queries{},
	}
}
