package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetAuthorByNameRequest struct {
	Name string
}

type getAuthorByNameHandler struct {
	repository repository.Repository
}

type GetAuthorByNameHandler decorator.QueryHandler[GetAuthorByNameRequest, *authors.Author]

func (h getAuthorByNameHandler) Handle(c context.Context, params GetAuthorByNameRequest) (*authors.Author, error) {
	author, err := h.repository.GetAuthorByName(c, params.Name)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return author, nil
}

func NewGetAuthorByNameHandler(repository repository.Repository, logger *zap.SugaredLogger) GetAuthorByNameHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getAuthorByNameHandler{
			repository: repository,
		}, logger,
	)
}
