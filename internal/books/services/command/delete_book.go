package command

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type DeleteBookRequest struct {
	ID uuid.UUID
}

type deleteBookHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
	cache      repository.Cache
}

type DeleteBookHandler decorator.CommandHandler[DeleteBookRequest]

func (h deleteBookHandler) Handle(c context.Context, params DeleteBookRequest) error {
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

func NewDeleteBookHandler(manager database.TransactionManager, repository repository.Repository, cache repository.Cache, logger *zap.SugaredLogger, metricsClient metrics.Client) DeleteBookHandler {
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
		}, logger, metricsClient,
	)
}
