package query

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type GetBooksParams struct {
	Title *string
}

type GetBooksHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

func (h GetBooksHandler) Execute(c context.Context, params GetBooksParams) ([]books.Book, error) {
	booksObj, err := h.cache.GetBooks(c, books.GetBooksDto{Title: params.Title})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	if booksObj != nil {
		return booksObj, nil
	}

	booksObj, err = h.repository.GetBooks(c, books.GetBooksDto{Title: params.Title})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	if booksObj == nil {
		return []books.Book{}, nil
	}

	err = h.cache.SetBooks(c, books.GetBooksDto{Title: params.Title}, booksObj)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return booksObj, nil
}

func NewGetBooksHandler(repository repository.Repository, cache repository.Cache) GetBooksHandler {
	return GetBooksHandler{repository, cache}
}
