package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	AuthorId  uuid.UUID `json:"authorId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
