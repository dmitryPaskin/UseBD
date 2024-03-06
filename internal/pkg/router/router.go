package router

import (
	_ "BDServer/docs"
	"BDServer/internal/controller"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwager "github.com/swaggo/http-swagger"
)

type Router struct {
	*chi.Mux
}

func New(userController *controller.UserController) Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwager.WrapHandler)

	r.Post("/api/users", userController.Create)
	r.Get("/api/users/{id}", userController.GetById)
	r.Put("/api/users/{id}", userController.Update)
	r.Delete("/api/users/{id}", userController.Delete)
	r.Get("/api/users", userController.List)

	return Router{r}
}
