package http

import (
	"go-boilerplate/config"
	authors "go-boilerplate/internal/author/ports"
	books "go-boilerplate/internal/book/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
}

func RunHTTPServer() {
	app := fiber.New()
	config, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	api := app.Group("/api")
	api.Route("/v1/authors", authors.RegisterAuthorHTTPRoutes)
	api.Route("/v1/books", books.RegisterBookHTTPRoutes)

	log.Fatal(app.Listen(":" + config.ServerPort))
}
