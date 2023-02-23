package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type RedisRepository interface {
	GetByID(ctx context.Context, key string) (*models.User, error)
	SetUser(ctx context.Context, userKey string, duration int, user *models.User) error
}
