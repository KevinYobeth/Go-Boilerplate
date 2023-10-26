package book

import (
	"context"
	"library/domain/author"
	"library/shared"
	model "library/shared/models"

	"github.com/google/uuid"
)

type GetAllBookReturn struct {
	Books []model.Book `json:"books"`
}

type Handler struct {
	UseCase UseCase
}

type UseCase struct {
	Repo          Repo
	AuthorUseCase author.UseCase
}

type Repo interface {
	Insert(ctx context.Context, book model.Book) error
	GetById(ctx context.Context, bookId uuid.UUID) (model.Book, error)
	DeleteById(ctx context.Context, bookId uuid.UUID) error
	Update(ctx context.Context, bookId uuid.UUID, book model.Book) error
	GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllBookReturn, error)
}

type UpsertBookEntity struct {
	Title    string    `json:"title"`
	AuthorId uuid.UUID `json:"authorId"`
}
