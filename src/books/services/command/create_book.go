package command

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type CreateBookParams struct {
	Title  string
	Author string
}

type CreateBookHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

func (h CreateBookHandler) Execute(c context.Context, params CreateBookParams) error {
	dto := books.NewCreateBookDto(params.Title)

	err := h.repository.CreateBook(c, dto)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return h.cache.ClearBooks(c)
}

func NewCreateBookHandler(database repository.Repository, cache repository.Cache) CreateBookHandler {
	return CreateBookHandler{database, cache}
}
