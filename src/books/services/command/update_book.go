package command

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type UpdateBookHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

func (h UpdateBookHandler) Execute(c context.Context, request books.UpdateBookDto) error {
	err := h.repository.UpdateBook(c, request)
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = h.cache.ClearBooks(c)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewUpdateBookHandler(repository repository.Repository, cache repository.Cache) UpdateBookHandler {
	return UpdateBookHandler{repository, cache}
}
