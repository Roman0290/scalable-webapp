package repository

import (
	"context"

	"example.com/go-project_1/backend/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	FindByEmail(ctx context.Context, email string) (models.User, error)
	FindByID(ctx context.Context, id string) (models.User, error)
}

type TokenRepository interface {
	StoreRefreshToken(ctx context.Context, userID, tokenHash string) error
	RevokeRefreshToken(ctx context.Context, tokenHash string) error
}