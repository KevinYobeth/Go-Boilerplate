package transport

import (
	"github.com/gofiber/fiber/v2"
)

type Server interface {
	RegisterBookHTTPRoutes(r fiber.Router)

	GetBooks(c *fiber.Ctx) error
	GetBook(c *fiber.Ctx) error

	CreateBook(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
}
