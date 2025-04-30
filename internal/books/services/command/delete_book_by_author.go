package command

import (
	"context"
	"go-boilerplate/internal/books/infrastructure/repository"
	"go-boilerplate/internal/books/services/helper"
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/metrics"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type DeleteBookByAuthorRequest struct {
	AuthorID uuid.UUID
}

type deleteBookByAuthorHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
}

type DeleteBookByAuthorHandler decorator.CommandHandler[DeleteBookByAuthorRequest]

func (h deleteBookByAuthorHandler) Handle(c context.Context, params DeleteBookByAuthorRequest) error {
	return tracerr.Wrap(h.manager.RunInTransaction(c, func(c context.Context) error {
		books, err := helper.GetBooksByAuthor(c, helper.GetBooksByAuthorOpts{
			Params: helper.GetBooksByAuthorRequest{
				ID: params.AuthorID,
			},
		})
		if err != nil {
			return tracerr.Wrap(err)
		}

		bookUUIDs := make([]uuid.UUID, len(books))
		for i, book := range books {
			bookUUIDs[i] = book.ID
		}

		err = h.repository.DeleteBooks(c, bookUUIDs)
		if err != nil {
			return tracerr.Wrap(err)
		}

		return nil
	}))
}

func NewDeleteBookByAuthorHandler(manager database.TransactionManager, repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) DeleteBookByAuthorHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyCommandDecorators(
		deleteBookByAuthorHandler{
			manager:    manager,
			repository: repository,
		}, logger, metricsClient,
	)
}
