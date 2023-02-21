package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
}
