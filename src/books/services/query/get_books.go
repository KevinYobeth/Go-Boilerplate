package query

import (
	"context"
	"fmt"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type GetBooksHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

func (h GetBooksHandler) Execute(c context.Context, request books.GetBooksDto) ([]books.Book, error) {
	books, err := h.cache.GetBooks(c, request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	if books != nil {
		fmt.Println("From cache")
		return books, nil
	}

	fmt.Println("From db")
	books, err = h.repository.GetBooks(c, request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	err = h.cache.SetBooks(c, request, books)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return books, nil
}

func NewGetBooksHandler(repository repository.Repository, cache repository.Cache) GetBooksHandler {
	return GetBooksHandler{repository, cache}
}
