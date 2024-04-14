package transport

import (
	respond "go-boilerplate/shared/response"
	"go-boilerplate/shared/types"
	"go-boilerplate/shared/utils"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app *services.Application
}

func NewHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) RegisterBookHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

// GET /books
func (h HTTPTransport) GetBooks(c echo.Context) error {
	booksObj, err := h.app.Queries.GetBooks.Execute(c.Request().Context(), books.GetBooksDto{})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, types.ResponseBody{
		Data:    booksObj,
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

	book, err := h.app.Queries.GetBook.Execute(c.Request().Context(), parsedUUID)
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, types.ResponseBody{
		Data:    book,
		Message: "success get book",
	})
	return nil
}

// POST /books
func (h HTTPTransport) CreateBook(c echo.Context) error {
	var request books.CreateBookDto
	if err := c.Bind(&request); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	if err := h.app.Commands.CreateBook.Execute(c.Request().Context(), request); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusCreated, types.ResponseBody{
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

	if err := h.app.Commands.UpdateBook.Execute(c.Request().Context(), books.UpdateBookDto{
		ID:    parsedUUID,
		Title: request.Title,
	}); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, types.ResponseBody{
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

	if err := h.app.Commands.DeleteBook.Execute(c.Request().Context(), parsedUUID); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, types.ResponseBody{
		Message: "success delete book",
	})
	return nil
}
