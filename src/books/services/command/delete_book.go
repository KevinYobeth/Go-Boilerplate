package command

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type DeleteBookParams struct {
	ID uuid.UUID
}

type deleteBookHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
	cache      repository.Cache
}

type DeleteBookHandler decorator.CommandHandler[DeleteBookParams]

func (h deleteBookHandler) Handle(c context.Context, params DeleteBookParams) error {
	return tracerr.Wrap(h.manager.RunInTransaction(c, func(c context.Context) error {
		bookObj, err := helper.GetBook(c, helper.GetBookOpts{
			Params: helper.GetBookRequest{
				ID: params.ID,
			},
			BookRepository: h.repository,
		})
		if err != nil {
			return tracerr.Wrap(err)
		}

		err = h.repository.DeleteBook(c, bookObj.ID)
		if err != nil {
			return tracerr.Wrap(err)
		}

		err = h.cache.ClearBooks(c)
		if err != nil {
			return tracerr.Wrap(err)
		}

		return nil
	}))
}

func NewDeleteBookHandler(manager database.TransactionManager, repository repository.Repository, cache repository.Cache, logger *zap.SugaredLogger) DeleteBookHandler {
	if repository == nil {
		panic("repository is required")
	}
	if cache == nil {
		panic("cache is required")
	}

	return decorator.ApplyCommandDecorators(
		deleteBookHandler{
			manager:    manager,
			repository: repository,
			cache:      cache,
		}, logger,
	)
}
