package book

import (
	"context"
	"library/domain/author"
	"library/shared"
	model "library/shared/models"

	"github.com/google/uuid"
)

type Handler struct {
	UseCase UseCase
}

type UseCase struct {
	Repo          Repo
	AuthorUseCase author.UseCase
}

type Repo interface {
	GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllBookReturn, error)
	GetById(ctx context.Context, bookId uuid.UUID) (model.Book, error)
	Create(ctx context.Context, book model.Book) error
	UpdateById(ctx context.Context, bookId uuid.UUID, book model.Book) error
	DeleteById(ctx context.Context, bookId uuid.UUID) error
}

type UpsertBookEntity struct {
	Title    string    `json:"title"`
	AuthorId uuid.UUID `json:"authorId"`
}

type GetAllBookReturn struct {
	Books []model.Book `json:"books"`
	Count int64        `json:"-"`
}
