package bookings

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateBookingDto struct {
	ID       uuid.UUID `json:"id"`
	BookID   uuid.UUID `json:"book_id"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

func NewCreateBookingDto(bookID uuid.UUID, dateFrom time.Time, dateTo time.Time) CreateBookingDto {
	return CreateBookingDto{
		ID:       uuid.New(),
		BookID:   bookID,
		DateFrom: dateFrom,
		DateTo:   dateTo,
	}
}

func (m *CreateBookingDto) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	fmt.Println("before create hooks", tx.Statement.Context)

	return nil
}
