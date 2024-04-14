package command

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type UpdateBookParams struct {
	ID    uuid.UUID
	Title string
}

type UpdateBookHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

func (h UpdateBookHandler) Execute(c context.Context, params UpdateBookParams) error {
	err := h.repository.UpdateBook(c, books.UpdateBookDto{
		ID:    params.ID,
		Title: params.Title,
	})
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
