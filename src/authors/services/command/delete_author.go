package command

import (
	"context"
	"go-boilerplate/shared/event"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type DeleteAuthorParams struct {
	ID uuid.UUID
}

type DeleteAuthorHandler struct {
	repository repository.Repository
	publisher  event.PublisherInterface
}

func (h DeleteAuthorHandler) Execute(c context.Context, params DeleteAuthorParams) error {
	err := h.repository.DeleteAuthor(c, params.ID)
	if err != nil {
		return tracerr.Wrap(err)
	}

	qName := "bobah"
	err = h.publisher.Publish(c, event.NewEvent("author.delete",
		event.PublisherOptions{Topic: "HELLO WORLD", Queue: &qName}),
	)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewDeleteAuthorHandler(database repository.Repository, publisher event.PublisherInterface) DeleteAuthorHandler {
	return DeleteAuthorHandler{database, publisher}
}
