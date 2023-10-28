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
	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/authors", a.loadAuthorRoutes)
	router.Route("/books", a.loadBookRoutes)

	a.router = router
}

func (a *App) loadAuthorRoutes(router chi.Router) {
	authorRepository := author.NewAuthorPostgresRepo(a.pgdb)
	authorUseCase := author.NewAuthorUseCase(authorRepository)
	authorHandler := author.NewAuthorHandler(*authorUseCase)

	router.Get("/", authorHandler.GetAll)
	router.Post("/", authorHandler.Create)
	router.Put("/{id}", authorHandler.UpdateById)
	router.Delete("/{id}", authorHandler.DeleteById)
	router.Get("/{id}", authorHandler.GetById)
}

func (a *App) loadBookRoutes(router chi.Router) {
	authorRepository := author.NewAuthorPostgresRepo(a.pgdb)
	authorUseCase := author.NewAuthorUseCase(authorRepository)

	bookRepository := book.NewBookPostgresRepo(a.pgdb)
	bookUseCase := book.NewBookUseCase(bookRepository, *authorUseCase)
	bookHandler := book.NewBookHandler(*bookUseCase)

	router.Get("/", bookHandler.GetAll)
	router.Post("/", bookHandler.Create)
	router.Put("/{id}", bookHandler.UpdateById)
	router.Delete("/{id}", bookHandler.DeleteById)
	router.Get("/{id}", bookHandler.GetById)
}
