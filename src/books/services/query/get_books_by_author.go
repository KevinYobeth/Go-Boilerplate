package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetBooksByAuthorParams struct {
	ID uuid.UUID
}

type getBooksByAuthorHandler struct {
	repository repository.Repository
}

type GetBooksByAuthorHandler decorator.QueryHandler[GetBooksByAuthorParams, []books.Book]

func (h getBooksByAuthorHandler) Handle(c context.Context, params GetBooksByAuthorParams) ([]books.Book, error) {
	booksObj, err := helper.GetBooksByAuthor(c, helper.GetBooksByAuthorOpts{
		Params: helper.GetBooksByAuthorRequest{
			ID: params.ID,
		},
		BookRepository: h.repository,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return booksObj, nil
}

func NewGetBooksByAuthorHandler(repository repository.Repository, logger *zap.SugaredLogger) GetBooksByAuthorHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getBooksByAuthorHandler{
			repository: repository,
		}, logger)
}
