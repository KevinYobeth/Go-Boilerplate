package command

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type DeleteBookParams struct {
	ID uuid.UUID
}

type DeleteBookHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
	cache      repository.Cache
}

func (h DeleteBookHandler) Execute(c context.Context, params DeleteBookParams) error {
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

func NewDeleteBookHandler(manager database.TransactionManager, repository repository.Repository, cache repository.Cache) DeleteBookHandler {
	return DeleteBookHandler{manager, repository, cache}
}
