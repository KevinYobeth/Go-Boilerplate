package author

import (
	"context"
	"library/shared"
)

type GetAllAuthorReturn struct {
	Authors []Author `json:"authors"`
}

type Repo interface {
	Insert(ctx context.Context, author Author) error
	GetById(ctx context.Context, authorId string) (Author, error)
	DeleteById(ctx context.Context, authorId string) error
	Update(ctx context.Context, authorId string, author Author) error
	GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error)
}

type UpsertAuthorEntity struct {
	Name string `json:"name"`
}
