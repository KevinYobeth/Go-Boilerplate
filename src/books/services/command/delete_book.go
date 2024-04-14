package command

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/query"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type DeleteBookService struct {
	GetBook query.GetBookHandler
}

type DeleteBookHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
	cache      repository.Cache
	services   DeleteBookService
}

func (h DeleteBookHandler) Execute(c context.Context, id uuid.UUID) error {
	return tracerr.Wrap(h.manager.RunInTransaction(c, func(c context.Context) error {
		bookObj, err := h.services.GetBook.Execute(c, id)
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

func NewDeleteBookHandler(manager database.TransactionManager, repository repository.Repository, cache repository.Cache, services DeleteBookService) DeleteBookHandler {
	return DeleteBookHandler{manager, repository, cache, services}
}
