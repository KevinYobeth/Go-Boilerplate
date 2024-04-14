package ports

import (
	"go-boilerplate/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RunHTTPServer() {
	app := fiber.New()
	config, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":" + config.ServerPort))
}
