package query

import (
	"context"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/bookings/domain/bookings"
	"go-boilerplate/src/bookings/infrastructure/repository"

	"github.com/google/uuid"
)

type GetBookingParams struct {
	ID uuid.UUID
}

type GetBookingHandler struct {
	repository repository.GormRepository
}

func (h GetBookingHandler) Execute(c context.Context, params GetBookingParams) (*bookings.Booking, error) {
	booking, err := h.repository.GetBooking(c, params.ID)
	if err != nil {
		return nil, err
	}

	if booking == nil {
		return nil, errors.NewNotFoundError(nil, "booking")
	}

	return booking, nil
}

func NewGetBookingHandler(repository repository.GormRepository) GetBookingHandler {
	return GetBookingHandler{repository}
}
