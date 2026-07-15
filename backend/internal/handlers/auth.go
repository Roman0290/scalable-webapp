package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"example.com/go-project_1/backend/internal/services"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload registerRequest
	if err := decodeJSON(r, &payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json payload"})
		return
	}

	if err := validateRegister(payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	result, err := h.service.Register(r.Context(), services.RegisterInput{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, result)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload loginRequest
	if err := decodeJSON(r, &payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json payload"})
		return
	}

	if err := validateLogin(payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	result, err := h.service.Login(r.Context(), services.LoginInput{
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var payload tokenRequest
	if err := decodeJSON(r, &payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json payload"})
		return
	}

	if strings.TrimSpace(payload.RefreshToken) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "refreshToken is required"})
		return
	}

	result, err := h.service.Refresh(r.Context(), payload.RefreshToken)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var payload tokenRequest
	if err := decodeJSON(r, &payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json payload"})
		return
	}

	if strings.TrimSpace(payload.RefreshToken) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "refreshToken is required"})
		return
	}

	if err := h.service.Logout(r.Context(), payload.RefreshToken); err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "logged out"})
}

func validateRegister(payload registerRequest) error {
	if strings.TrimSpace(payload.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(payload.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(payload.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

func validateLogin(payload loginRequest) error {
	if strings.TrimSpace(payload.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(payload.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

func decodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}

func writeServiceError(w http.ResponseWriter, err error) {
	if errors.Is(err, services.ErrNotImplemented) {
		writeJSON(w, http.StatusNotImplemented, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}