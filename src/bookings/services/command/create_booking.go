package command

import (
	"context"
	"fmt"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/bookings/domain/bookings"
	"go-boilerplate/src/bookings/infrastructure/repository"
	bookService "go-boilerplate/src/books/services"
	bookQuery "go-boilerplate/src/books/services/query"
	"time"

	"github.com/google/uuid"
)

type CreateBookingParams struct {
	BookID   uuid.UUID
	DateFrom time.Time
	DateTo   time.Time
}

type CreateBookingService struct {
	BookService bookService.Application
}

type CreateBookingHandler struct {
	repository repository.GormRepository
	manager    database.GormTransactionManager
	service    CreateBookingService
}

func (h CreateBookingHandler) Execute(c context.Context, params CreateBookingParams) error {
	return h.manager.RunInTransaction(c, func(c context.Context) error {
		fmt.Println("tx running in transaction")
		dto := bookings.NewCreateBookingDto(uuid.MustParse("53521371-81d8-4211-824a-c16d28d14e77"), params.DateFrom, params.DateTo)

		err := h.repository.CreateBooking(c, dto)
		fmt.Println("tx create booking")
		if err != nil {
			return err
		}

		_, err = h.service.BookService.Queries.GetBook.Execute(c, bookQuery.GetBookParams{
			ID: params.BookID,
		})
		fmt.Println("tx get book")
		if err != nil {
			fmt.Println("tx get book err", err)
			return err
		}

		return nil
	})
}

func NewCreateBookingHandler(repository repository.GormRepository, manager database.GormTransactionManager, service CreateBookingService) CreateBookingHandler {
	return CreateBookingHandler{repository, manager, service}
}
