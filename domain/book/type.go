package book

import (
	"context"
	helper "library/utils"
)

type GetAllBookReturn struct {
	Books  []Book
	Cursor uint64
}

type Repo interface {
	Insert(ctx context.Context, book Book) error
	GetById(ctx context.Context, bookId string) (Book, error)
	DeleteById(ctx context.Context, bookId string) error
	Update(ctx context.Context, bookId string, book Book) error
	GetAll(ctx context.Context, pagination helper.LimitPagination) (GetAllBookReturn, error)
}
