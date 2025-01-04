package query

import (
	"context"
	"go-boilerplate/src/bookings/domain/bookings"
	"go-boilerplate/src/bookings/infrastructure/repository"
)

type GetBookingsParams struct {
}

type GetBookingsHandler struct {
	repository repository.GormRepository
}

func (h GetBookingsHandler) Execute(c context.Context) ([]bookings.Booking, error) {
	bookingsObj, err := h.repository.GetBookings(c)
	if err != nil {
		return nil, err
	}

	return bookingsObj, nil
}

func NewGetBookingsHandler(repository repository.GormRepository) GetBookingsHandler {
	return GetBookingsHandler{repository}
}
