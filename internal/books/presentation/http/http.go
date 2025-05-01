package http

import (
	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/books"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/query"
	response "github.com/kevinyobeth/go-boilerplate/shared/response"
	"github.com/kevinyobeth/go-boilerplate/shared/types"
	"github.com/kevinyobeth/go-boilerplate/shared/utils"

	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app *services.Application
}

func NewBooksHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

// GET /books
func (h HTTPTransport) GetBooks(c echo.Context) error {
	booksObj, err := h.app.Queries.GetBooks.Handle(c.Request().Context(), query.GetBooksRequest{})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: GetBooksResponse{
			Data:    TransformToHTTPBooksWithAuthor(booksObj),
			Message: "success get books",
		},
	})
	return nil
}

// GET /books/:id
func (h HTTPTransport) GetBook(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	book, err := h.app.Queries.GetBook.Handle(c.Request().Context(), query.GetBookRequest{ID: parsedUUID})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: GetBookResponse{
			Data:    TransformToHTTPBook(book),
			Message: "success get book",
		},
	})
	return nil
}

// POST /books
func (h HTTPTransport) CreateBook(c echo.Context) error {
	var request CreateBookRequest
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	if err := h.app.Commands.CreateBook.Handle(c.Request().Context(),
		command.CreateBookRequest{Title: request.Title, Author: request.Author}); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: MessageResponse{
			Message: "success create book",
		},
	})
	return nil
}

// PUT /books/:id
func (h HTTPTransport) UpdateBook(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return nil
	}

	var request books.UpdateBookDto
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	if err := h.app.Commands.UpdateBook.Handle(c.Request().Context(), command.UpdateBookRequest{
		ID:    parsedUUID,
		Title: request.Title,
	}); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: MessageResponse{
			Message: "success update book",
		},
	})
	return nil
}

// DELETE /books/:id
func (h HTTPTransport) DeleteBook(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return nil
	}

	if err := h.app.Commands.DeleteBook.Handle(c.Request().Context(), command.DeleteBookRequest{
		ID: parsedUUID,
	}); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: MessageResponse{
			Message: "success delete book",
		},
	})
	return nil
}
