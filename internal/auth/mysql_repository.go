package auth

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

type Repository interface {
	Register(ctx context.Context, user *models.UserCreateRequest) (*db.User, error)
	FindByEmail(ctx context.Context, email string) (*db.User, error)
	GetUsers(ctx context.Context, req *utils.PaginationRequest) ([]*db.User, error)
	GetCount(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, userID string) (*db.User, error)
}
