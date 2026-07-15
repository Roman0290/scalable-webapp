package middleware

import (
	"net/http"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func Chain(next http.Handler) http.Handler {
	handler := chiMiddleware.RequestID(next)
	handler = chiMiddleware.RealIP(handler)
	handler = chiMiddleware.Logger(handler)
	handler = chiMiddleware.Recoverer(handler)
	handler = timeout(handler, 30*time.Second)
	return handler
}

func timeout(next http.Handler, duration time.Duration) http.Handler {
	return http.TimeoutHandler(next, duration, `{"error":"request timeout"}`)
}