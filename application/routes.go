package application

import (
	"library/domain/author"
	"library/domain/book"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/authors", a.loadAuthorRoutes)
	router.Route("/books", a.loadBookRoutes)

	a.router = router
}

func (a *App) loadAuthorRoutes(router chi.Router) {
	authorHandler := &author.AuthorInterface{
		Repo: &author.PostgresRepo{
			Client: a.pgdb,
		},
	}

	router.Put("/{id}", authorHandler.Update)
	router.Delete("/{id}", authorHandler.DeleteById)
	router.Get("/{id}", authorHandler.GetById)
	router.Get("/", authorHandler.List)
	router.Post("/", authorHandler.Create)
}

func (a *App) loadBookRoutes(router chi.Router) {
	bookHandler := &book.BookInterface{
		Repo: &book.PostgresRepo{
			Client: a.pgdb,
		},
	}

	router.Put("/{id}", bookHandler.Update)
	router.Delete("/{id}", bookHandler.DeleteById)
	router.Get("/{id}", bookHandler.GetById)
	router.Get("/", bookHandler.List)
	router.Post("/", bookHandler.Create)
}
