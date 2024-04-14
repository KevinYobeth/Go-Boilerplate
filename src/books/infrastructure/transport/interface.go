package transport

import (
	"github.com/labstack/echo/v4"
)

type HttpServer interface {
	RegisterBookHTTPRoutes(r *echo.Group)

	GetBooks(c echo.Context) error
	GetBook(c echo.Context) error

	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}
