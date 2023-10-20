package author

import (
	"context"
	helper "library/utils"
)

type GetAllAuthorReturn struct {
	Authors []Author `json:"authors"`
	Cursor  uint64   `json:"cursor"`
}

type Repo interface {
	Insert(ctx context.Context, author Author) error
	GetById(ctx context.Context, authorId string) (Author, error)
	DeleteById(ctx context.Context, authorId string) error
	Update(ctx context.Context, authorId string, author Author) error
	GetAll(ctx context.Context, pagination helper.LimitPagination) (GetAllAuthorReturn, error)
}

type CreateAuthorEntity struct {
	Name string `json:"name"`
}
