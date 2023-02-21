package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type Repository interface {
	Register(ctx context.Context, user *models.UserCreateRequest) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
