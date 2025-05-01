package command

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/authors"
	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/books"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/intraprocess"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"

	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type CreateBookRequest struct {
	Title  string
	Author string
}

type createBookHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
	cache      repository.Cache

	authorService intraprocess.BookAuthorIntraprocess
}

type CreateBookHandler decorator.CommandHandler[CreateBookRequest]

func (h createBookHandler) Handle(c context.Context, params CreateBookRequest) error {
	return tracerr.Wrap(h.manager.RunInTransaction(c, func(c context.Context) error {
		var authorObj *authors.Author

		authorObj, err := h.authorService.GetAuthorByName(c, params.Author)
		if err != nil {
			return tracerr.Wrap(err)
		}
		if authorObj == nil {
			authorObj, err = h.authorService.CreateAuthor(c, params.Author)
			if err != nil {
				return tracerr.Wrap(err)
			}
		}

		dto := books.NewCreateBookDto(params.Title)

		err = h.repository.CreateBook(c, dto)
		if err != nil {
			return tracerr.Wrap(err)
		}

		err = helper.CreateAuthorBook(c, helper.CreateAuthorBookOpts{
			Params: helper.CreateAuthorBookRequest{
				BookID:   dto.ID,
				AuthorID: authorObj.ID,
			},
			BookRepository: h.repository,
		})
		if err != nil {
			return tracerr.Wrap(err)
		}

		return h.cache.ClearBooks(c)
	}))
}

func NewCreateBookHandler(manager database.TransactionManager, database repository.Repository, cache repository.Cache, authorService intraprocess.BookAuthorIntraprocess, logger *zap.SugaredLogger, metricsClient metrics.Client) CreateBookHandler {
	if database == nil {
		panic("nil database")
	}
	if cache == nil {
		panic("nil cache")
	}
	if authorService == nil {
		panic("nil authorService")
	}

	return decorator.ApplyCommandDecorators(
		createBookHandler{
			manager:    manager,
			repository: database,
			cache:      cache,

			authorService: authorService,
		}, logger, metricsClient,
	)
}
