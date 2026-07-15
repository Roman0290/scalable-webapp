package routes

import (
	"net/http"

	"example.com/go-project_1/backend/internal/handlers"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	Health *handlers.HealthHandler
	Auth   *handlers.AuthHandler
}

func New(deps Dependencies) http.Handler {
	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", deps.Health.Handle)
		r.Post("/auth/register", deps.Auth.Register)
		r.Post("/auth/login", deps.Auth.Login)
		r.Post("/auth/refresh", deps.Auth.Refresh)
		r.Post("/auth/logout", deps.Auth.Logout)
	})

	return router
}