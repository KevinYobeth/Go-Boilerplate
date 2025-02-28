package authors

import (
	"context"
)

type AuthorService interface {
	GetAuthors(c context.Context, name string) ([]Author, error)
	CreateAuthor(c context.Context, name string) (*Author, error)
}
