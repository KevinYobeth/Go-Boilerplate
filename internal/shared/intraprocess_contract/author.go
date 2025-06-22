package intraprocess_contract

import (
	"context"

	"github.com/google/uuid"
)

type Author struct {
	ID   uuid.UUID `json:"author_id"`
	Name string    `json:"author_name"`
}

type AuthorInterface interface {
	GetAuthors(c context.Context, name *string) ([]Author, error)
	CreateAuthor(c context.Context, name string) (*Author, error)
}
