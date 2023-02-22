package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

type UseCase interface {
	Register(ctx context.Context, user *models.UserCreateRequest) (*models.UserWithToken, error)
	Login(ctx context.Context, user *models.UserLoginRequest) (*models.UserWithToken, error)
	GetUsers(ctx context.Context, req *utils.PaginationRequest) (*models.UsersList, error)
}
