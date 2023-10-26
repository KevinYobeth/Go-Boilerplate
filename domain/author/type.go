package author

import (
	"context"
	"library/shared"
	model "library/shared/models"

	"github.com/google/uuid"
)

type GetAllAuthorReturn struct {
	Authors []model.Author `json:"authors"`
	Count   int64          `json:"-"`
}

type Handler struct {
	UseCase UseCase
}

type UseCase struct {
	Repo Repo
}

type Repo interface {
	GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error)
	GetById(ctx context.Context, authorId uuid.UUID) (model.Author, error)
	Create(ctx context.Context, author model.Author) error
	UpdateById(ctx context.Context, authorId uuid.UUID, author model.Author) error
	DeleteById(ctx context.Context, authorId uuid.UUID) error
}

type UpsertAuthorEntity struct {
	Name string `json:"name"`
}
