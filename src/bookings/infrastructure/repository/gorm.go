package repository

import (
	"context"
	"fmt"
	"go-boilerplate/src/bookings/domain/bookings"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type GormBookingsRepo struct {
	db *gorm.DB
}

func NewBookingsGormRepository(db *gorm.DB) GormRepository {
	return &GormBookingsRepo{db}
}

func (r GormBookingsRepo) GetBookings(c context.Context) ([]bookings.Booking, error) {
	bookings := []bookings.Booking{}

	err := r.db.Find(&bookings).Error
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return bookings, nil
}

func (r GormBookingsRepo) CreateBooking(c context.Context, request bookings.CreateBookingDto) error {
	fmt.Println("IN CREATE BOOKINg", r.db.Statement.Context)
	err := r.db.Table("bookings").Create(&request).Error
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (r GormBookingsRepo) GetBooking(c context.Context, id uuid.UUID) (*bookings.Booking, error) {
	booking := &bookings.Booking{}

	err := r.db.Where("id = ?", id).First(&booking).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, tracerr.Wrap(err)
	}

	return booking, nil
}
