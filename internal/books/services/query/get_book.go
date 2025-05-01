package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/books"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetBookRequest struct {
	ID uuid.UUID
}

type getBookHandler struct {
	repository repository.Repository
}

type GetBookHandler decorator.QueryHandler[GetBookRequest, *books.Book]

func (h getBookHandler) Handle(c context.Context, params GetBookRequest) (*books.Book, error) {
	book, err := helper.GetBook(c, helper.GetBookOpts{
		Params: helper.GetBookRequest{
			ID: params.ID,
		},
		BookRepository: h.repository,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return book, nil
}

func NewGetBookHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetBookHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getBookHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
