package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetBooksRequest struct {
	Title *string
}

type getBooksHandler struct {
	repository repository.Repository
	cache      repository.Cache
}

type GetBooksHandler decorator.QueryHandler[GetBooksRequest, []books.BookWithAuthor]

func (h getBooksHandler) Handle(c context.Context, params GetBooksRequest) ([]books.BookWithAuthor, error) {
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
		return []books.BookWithAuthor{}, nil
	}

	err = h.cache.SetBooks(c, books.GetBooksDto{Title: params.Title}, booksObj)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return booksObj, nil
}

func NewGetBooksHandler(repository repository.Repository, cache repository.Cache, logger *zap.SugaredLogger) GetBooksHandler {
	if repository == nil {
		panic("repository is required")
	}
	if cache == nil {
		panic("cache is required")
	}

	return decorator.ApplyQueryDecorators(
		getBooksHandler{
			repository: repository,
			cache:      cache,
		}, logger)
}
