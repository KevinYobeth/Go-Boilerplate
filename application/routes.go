package application

import (
	"library/domain/author"
	"library/domain/order"
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

	router.Route("/orders", a.loadOrderRoutes)
	router.Route("/authors", a.loadAuthorRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	// orderHandler := &order.OrderInterface{
	// 	Repo: &order.RedisRepo{
	// 		Client: a.rdb,
	// 	},
	// }
	orderHandler := &order.OrderInterface{
		Repo: &order.PostgresRepo{},
	}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateById)
	router.Delete("/{id}", orderHandler.DeleteById)
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
