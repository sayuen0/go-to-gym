package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, user *models.UserCreateRequest) (*models.UserWithToken, error)
	Login(ctx context.Context, user *models.UserLoginRequest) (*models.UserWithToken, error)
}
