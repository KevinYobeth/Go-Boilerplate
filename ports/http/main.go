package http

import (
	"go-boilerplate/config"
	books "go-boilerplate/src/books/infrastructure/transport"
	"go-boilerplate/src/books/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunHTTPServer() {
	app := echo.New()
	app.Debug = true

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Pre(middleware.RemoveTrailingSlash())

	config := config.LoadServerConfig()

	app.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	booksService := services.NewBookService()
	booksServer := books.NewHTTPServer(&booksService)

	api := app.Group("/api")

	booksServer.RegisterBookHTTPRoutes(api)

	app.Logger.Fatal(app.Start(":" + config.ServerPort))
}
