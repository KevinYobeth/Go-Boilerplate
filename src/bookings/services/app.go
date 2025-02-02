package services

import (
	"go-boilerplate/shared/database"
	"go-boilerplate/src/bookings/infrastructure/repository"
	"go-boilerplate/src/bookings/services/command"
	"go-boilerplate/src/bookings/services/query"
	"go-boilerplate/src/books/services"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateBooking command.CreateBookingHandler
}

type Queries struct {
	GetBookings query.GetBookingsHandler
	GetBooking  query.GetBookingHandler
}

func NewBookingService() Application {
	db := database.InitGorm()
	repo := repository.NewBookingsGormRepository(db)

	gormManager := database.NewGormTransactionManager(db)
	bookService := services.NewBookService()

	return Application{
		Commands: Commands{
			CreateBooking: command.NewCreateBookingHandler(repo, gormManager, command.CreateBookingService{
				BookService: bookService,
			}),
		},
		Queries: Queries{
			GetBookings: query.NewGetBookingsHandler(repo),
			GetBooking:  query.NewGetBookingHandler(repo),
		},
	}
}
