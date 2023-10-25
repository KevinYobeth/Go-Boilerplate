package author

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Books     []Book    `json:"books"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Book struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	AuthorId  uuid.UUID `json:"authorId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
