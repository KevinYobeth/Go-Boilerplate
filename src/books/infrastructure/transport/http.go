package transport

import (
	"go-boilerplate/shared/utils"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	app *services.Application
}

func NewHTTPServer(app *services.Application) Server {
	return &HTTPServer{app: app}
}

func (h HTTPServer) RegisterBookHTTPRoutes(r fiber.Router) {
	r.Get("/", h.GetBooks)
	r.Get("/:id", h.GetBook)

	r.Post("/", h.CreateBook)
	r.Put("/:id", h.UpdateBook)
	r.Delete("/:id", h.DeleteBook)
}

// GET /books
func (h HTTPServer) GetBooks(c *fiber.Ctx) error {
	books, err := h.app.Queries.GetBooks.Execute(c, books.GetBooksDto{})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(books)
	return nil
}

// GET /book
func (h HTTPServer) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedUUID := utils.ParseUUID(c, id)

	book, err := h.app.Queries.GetBook.Execute(c, parsedUUID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(book)
	return nil
}

// POST /books
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

// PUT /books/:id
func (h HTTPServer) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedUUID := utils.ParseUUID(c, id)

	var request books.UpdateBookDto
	if err := c.BodyParser(&request); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	if err := h.app.Commands.UpdateBook.Execute(c, books.UpdateBookDto{
		ID:    parsedUUID,
		Title: request.Title,
	}); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.Status(fiber.StatusOK)
	return nil
}

// DELETE /books/:id
func (h HTTPServer) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedUUID := utils.ParseUUID(c, id)

	if err := h.app.Commands.DeleteBook.Execute(c, parsedUUID); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.Status(fiber.StatusNoContent)
	return nil
}
