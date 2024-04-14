package adapters

import "github.com/gofiber/fiber/v2"

func RegisterAuthorHTTPRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("yang ini author")
	})
}
