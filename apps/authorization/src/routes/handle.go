package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handlers(s *HandlerServices) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api/external", func(r chi.Router) {
		r.Post("/", createUser(s))
	})

	return r

}
