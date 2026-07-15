package handlers

import (
	"net/http"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":    "ok",
		"requestId": chiMiddleware.GetReqID(r.Context()),
	})
}