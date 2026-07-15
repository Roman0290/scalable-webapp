package services

import (
	"context"
	"errors"

	"example.com/go-project_1/backend/internal/models"
	"example.com/go-project_1/backend/internal/repository"
)

var ErrNotImplemented = errors.New("auth service not implemented yet")

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type AuthResult struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
}

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResult, error)
	Login(ctx context.Context, input LoginInput) (AuthResult, error)
	Refresh(ctx context.Context, refreshToken string) (AuthResult, error)
	Logout(ctx context.Context, refreshToken string) error
}

type authService struct {
	users  repository.UserRepository
	tokens repository.TokenRepository
}

func NewAuthService(users repository.UserRepository, tokens repository.TokenRepository) AuthService {
	return &authService{users: users, tokens: tokens}
}

func (s *authService) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	return AuthResult{}, ErrNotImplemented
}

func (s *authService) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	return AuthResult{}, ErrNotImplemented
}

func (s *authService) Refresh(ctx context.Context, refreshToken string) (AuthResult, error) {
	return AuthResult{}, ErrNotImplemented
}

func (s *authService) Logout(ctx context.Context, refreshToken string) error {
	return ErrNotImplemented
}