package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type Author struct {
	ID   uuid.UUID `json:"author_id"`
	Name string    `json:"author_name"`
}

type AuthorIntraprocess interface {
	GetAuthors(c context.Context, name *string) ([]Author, error)
	CreateAuthor(c context.Context, name string) (*Author, error)
}
