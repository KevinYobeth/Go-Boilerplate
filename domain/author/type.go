package author

import (
	"context"
	"library/shared"
	model "library/shared/models"
)

type GetAllAuthorReturn struct {
	Authors []model.Author `json:"authors"`
}

type Repo interface {
	Insert(ctx context.Context, author model.Author) error
	GetById(ctx context.Context, authorId string) (model.Author, error)
	DeleteById(ctx context.Context, authorId string) error
	Update(ctx context.Context, authorId string, author model.Author) error
	GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error)
}

type UpsertAuthorEntity struct {
	Name string `json:"name"`
}
