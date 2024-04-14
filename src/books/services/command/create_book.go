package command

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type CreateBookHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

func (h CreateBookHandler) Execute(c context.Context, request books.CreateBookDto) error {
	dto := books.NewCreateBookDto(request.Title)

	err := h.repository.CreateBook(c, dto)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return h.cache.ClearBooks(c)
}

func NewCreateBookHandler(database repository.Repository, cache repository.Cache) CreateBookHandler {
	return CreateBookHandler{database, cache}
}
