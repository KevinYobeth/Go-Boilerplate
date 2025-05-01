package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/books"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetBooksByAuthorRequest struct {
	ID uuid.UUID
}

type getBooksByAuthorHandler struct {
	repository repository.Repository
}

type GetBooksByAuthorHandler decorator.QueryHandler[GetBooksByAuthorRequest, []books.Book]

func (h getBooksByAuthorHandler) Handle(c context.Context, params GetBooksByAuthorRequest) ([]books.Book, error) {
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

func NewGetBooksByAuthorHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetBooksByAuthorHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getBooksByAuthorHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
