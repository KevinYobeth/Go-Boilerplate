package adapters

import (
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	app *services.Application
}

func NewHTTPServer(app *services.Application) *HTTPServer {
	return &HTTPServer{app: app}
}

func (h HTTPServer) RegisterBookHTTPRoutes(r fiber.Router) {
	r.Get("/", h.GetBooks)
	r.Post("/", h.CreateBook)
}

func (h HTTPServer) GetBooks(c *fiber.Ctx) error {
	books, err := h.app.Queries.GetBooks.Execute(c, books.GetBooksDto{})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(books)
	return nil
}

func (h HTTPServer) CreateBook(c *fiber.Ctx) error {
	var request books.CreateBookDto
	if err := c.BodyParser(&request); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	if err := h.app.Commands.CreateBook.Execute(c, request); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.Status(fiber.StatusCreated)
	return nil
}
