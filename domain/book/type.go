package book

import (
	"context"
	"library/shared"

	"github.com/google/uuid"
)

type GetAllBookReturn struct {
	Books []Book `json:"books"`
}

type Repo interface {
	Insert(ctx context.Context, book Book) error
	GetById(ctx context.Context, bookId string) (Book, error)
	DeleteById(ctx context.Context, bookId string) error
	Update(ctx context.Context, bookId string, book Book) error
	GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllBookReturn, error)
}

type UpsertBookEntity struct {
	Title    string    `json:"title"`
	AuthorId uuid.UUID `json:"authorId"`
}
