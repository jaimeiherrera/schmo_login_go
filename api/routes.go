package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(router *chi.Mux, handlers *Handlers) {
	// Global middlewares (applied to all routes)
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Routes
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", handlers.Ping)

		r.Route("/users", func(r chi.Router) {
			r.Post("/", handlers.CreateUser)
		})
	})
}
