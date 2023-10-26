package model

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Books     *[]Book   `json:"books,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
