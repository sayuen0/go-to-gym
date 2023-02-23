package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

// UseCase defines the interface for authentication related use cases.
type UseCase interface {
	// Register registers a new user.
	Register(ctx context.Context, user *models.UserCreateRequest) (*models.UserWithToken, error)
	// Login logs in a user.
	Login(ctx context.Context, user *models.UserLoginRequest) (*models.UserWithToken, error)
	// GetUsers returns a list of users.
	GetUsers(ctx context.Context, req *utils.PaginationRequest) (*models.UsersList, error)
	// GetByID returns a user by UUID.
	GetByID(ctx context.Context, uuid string) (*models.User, error)
}
