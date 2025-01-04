package repository

import (
	"context"
	"go-boilerplate/src/bookings/domain/bookings"

	"github.com/google/uuid"
)

type GormRepository interface {
	GetBookings(c context.Context) ([]bookings.Booking, error)
	GetBooking(c context.Context, id uuid.UUID) (*bookings.Booking, error)

	CreateBooking(c context.Context, request bookings.CreateBookingDto) error
}
