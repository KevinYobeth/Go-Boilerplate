package command

import (
	"context"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/query"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type DeleteBookService struct {
	GetBook query.GetBookHandler
}

type DeleteBookHandler struct {
	repository repository.Repository
	services   DeleteBookService
}

func (h DeleteBookHandler) Execute(c context.Context, id uuid.UUID) error {
	bookObj, err := h.services.GetBook.Execute(c, id)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return h.repository.DeleteBook(c, bookObj.ID)
}

func NewDeleteBookHandler(repository repository.Repository, services DeleteBookService) DeleteBookHandler {

	return DeleteBookHandler{repository, services}
}
