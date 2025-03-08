package command

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"
	"go-boilerplate/src/books/services/query"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type DeleteBookByAuthorParams struct {
	AuthorID uuid.UUID
}

type DeleteBookByAuthorService struct {
	GetBooksByAuthor query.GetBooksByAuthorHandler
}

type DeleteBookByAuthorHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
}

func (h DeleteBookByAuthorHandler) Execute(c context.Context, params DeleteBookByAuthorParams) error {
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

func NewDeleteBookByAuthorHandler(manager database.TransactionManager, repository repository.Repository) DeleteBookByAuthorHandler {
	return DeleteBookByAuthorHandler{manager, repository}
}
