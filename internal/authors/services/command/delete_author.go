package command

import (
	"context"
	"go-boilerplate/internal/authors/domain/authors"
	"go-boilerplate/internal/authors/infrastructure/repository"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/event"
	"go-boilerplate/shared/metrics"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type DeleteAuthorRequest struct {
	ID uuid.UUID
}

type deleteAuthorHandler struct {
	repository repository.Repository
	publisher  event.PublisherInterface
}

type DeleteAuthorHandler decorator.CommandHandler[DeleteAuthorRequest]

func (h deleteAuthorHandler) Handle(c context.Context, params DeleteAuthorRequest) error {
	err := h.repository.DeleteAuthor(c, params.ID)
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = h.publisher.Publish(c, event.NewEvent("author.delete", authors.DeleteAuthorEvent{ID: params.ID}))
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewDeleteAuthorHandler(database repository.Repository, publisher event.PublisherInterface, logger *zap.SugaredLogger, metricsClient metrics.Client) DeleteAuthorHandler {
	if database == nil {
		panic("database is required")
	}

	if publisher == nil {
		panic("publisher is required")
	}

	return decorator.ApplyCommandDecorators(
		deleteAuthorHandler{
			repository: database,
			publisher:  publisher,
		}, logger, metricsClient,
	)
}
