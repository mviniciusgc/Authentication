package routes

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Handlers(s *HandlerServices) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api/internal", func(r chi.Router) {
		r.Post("/", createUser(s))
	})

	r.Route("/api/external", func(r chi.Router) {
		r.Post("/authenticate", authenticate(s))
		r.Post("/refresh", refreshToken(s))
	})

	return r

}
