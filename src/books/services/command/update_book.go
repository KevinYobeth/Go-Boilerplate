package command

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type UpdateBookRequest struct {
	ID    uuid.UUID
	Title string
}

type updateBookHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

type UpdateBookHandler decorator.CommandHandler[UpdateBookRequest]

func (h updateBookHandler) Handle(c context.Context, params UpdateBookRequest) error {
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

func NewUpdateBookHandler(repository repository.Repository, cache repository.Cache, logger *zap.SugaredLogger) UpdateBookHandler {
	if repository == nil {
		panic("repository is required")
	}
	if cache == nil {
		panic("cache is required")
	}

	return decorator.ApplyCommandDecorators(
		updateBookHandler{
			repository: repository,
			cache:      cache,
		}, logger,
	)
}
