package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
)

type Repository interface {
	Register(ctx context.Context, user *models.UserCreateRequest) (*db.User, error)
	FindByEmail(ctx context.Context, email string) (*db.User, error)
}
