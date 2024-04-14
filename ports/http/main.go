package http

import (
	"go-boilerplate/config"
	authors "go-boilerplate/src/authors/infrastructure/adapters"
	books "go-boilerplate/src/books/infrastructure/adapters"
	"go-boilerplate/src/books/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RunHTTPServer() {
	app := fiber.New()
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New())

	config, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	booksService := services.NewBookService()
	booksServer := books.NewHTTPServer(&booksService)

	api := app.Group("/api")
	api.Route("/v1/authors", authors.RegisterAuthorHTTPRoutes)
	api.Route("/v1/books", booksServer.RegisterBookHTTPRoutes)

	log.Fatal(app.Listen(":" + config.ServerPort))
}
