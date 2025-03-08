package http

import (
	respond "go-boilerplate/shared/response"
	"go-boilerplate/shared/utils"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/services"
	"go-boilerplate/src/books/services/command"
	"go-boilerplate/src/books/services/query"
	"net/http"

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
	booksObj, err := h.app.Queries.GetBooks.Handle(c.Request().Context(), query.GetBooksParams{})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, GetBooksResponse{
		Data:    TransformToHTTPBooksWithAuthor(booksObj),
		Message: "success get books",
	})
	return nil
}

// GET /books/:id
func (h HTTPTransport) GetBook(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	book, err := h.app.Queries.GetBook.Handle(c.Request().Context(), query.GetBookParams{ID: parsedUUID})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, GetBookResponse{
		Data:    TransformToHTTPBook(book),
		Message: "success get book",
	})
	return nil
}

// POST /books
func (h HTTPTransport) CreateBook(c echo.Context) error {
	var request CreateBookRequest
	if err := c.Bind(&request); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	if err := h.app.Commands.CreateBook.Handle(c.Request().Context(),
		command.CreateBookParams{Title: request.Title, Author: request.Author}); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusCreated, MessageResponse{
		Message: "success create book",
	})
	return nil
}

// PUT /books/:id
func (h HTTPTransport) UpdateBook(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		respond.SendHTTP(c, err)
		return nil
	}

	var request books.UpdateBookDto
	if err := c.Bind(&request); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	if err := h.app.Commands.UpdateBook.Handle(c.Request().Context(), command.UpdateBookParams{
		ID:    parsedUUID,
		Title: request.Title,
	}); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, MessageResponse{
		Message: "success update book",
	})
	return nil
}

// DELETE /books/:id
func (h HTTPTransport) DeleteBook(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		respond.SendHTTP(c, err)
		return nil
	}

	if err := h.app.Commands.DeleteBook.Handle(c.Request().Context(), command.DeleteBookParams{
		ID: parsedUUID,
	}); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, MessageResponse{
		Message: "success delete book",
	})
	return nil
}
