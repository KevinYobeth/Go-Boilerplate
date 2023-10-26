package author

import (
	"context"
	"library/shared"
	model "library/shared/models"

	"github.com/google/uuid"
)

type GetAllAuthorReturn struct {
	Authors []model.Author `json:"authors"`
}

type Handler struct {
	UseCase UseCase
}

type UseCase struct {
	Repo Repo
}

type Repo interface {
	Insert(ctx context.Context, author model.Author) error
	GetById(ctx context.Context, authorId uuid.UUID) (model.Author, error)
	DeleteById(ctx context.Context, authorId uuid.UUID) error
	Update(ctx context.Context, authorId uuid.UUID, author model.Author) error
	GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error)
}

type UpsertAuthorEntity struct {
	Name string `json:"name"`
}
