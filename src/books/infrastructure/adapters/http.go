package adapters

import "github.com/gofiber/fiber/v2"

func RegisterBookHTTPRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hehehhe bisa nih")
	})
}
