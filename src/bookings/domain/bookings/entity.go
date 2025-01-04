package bookings

import (
	"go-boilerplate/src/books/domain/books"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID       uuid.UUID `json:"id"`
	BookID   uuid.UUID `json:"book_id"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
	Book     books.Book
}
