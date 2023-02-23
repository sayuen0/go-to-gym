package repository

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
	"github.com/sayuen0/go-to-gym/pkg/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type authRepo struct {
	db *sql.DB
}

// NewAuthRepo is a constructor of auth.Repository
func NewAuthRepo(db *sql.DB) auth.Repository {
	return &authRepo{db: db}
}

// Register inserts a new user record
func (r *authRepo) Register(ctx context.Context, user *models.UserCreateRequest) (*db.User, error) {
	u := &db.User{
		UserID:         utils.NewUUIDStr(),
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}

	if err := u.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "authRepo.Register.Insert")
	}

	return u, nil
}

// FindByEmail return a user by email
func (r *authRepo) FindByEmail(ctx context.Context, email string) (*db.User, error) {
	return db.Users(qm.Where("email=?", email)).One(ctx, r.db)
}

// GetUsers returns a list of users
func (r *authRepo) GetUsers(ctx context.Context, req *utils.PaginationRequest) ([]*db.User, error) {
	return db.Users(
		req.GenerateQueryMods()...,
	).All(ctx, r.db)
}

// GetCount returns the count of rows
func (r *authRepo) GetCount(ctx context.Context) (int64, error) {
	return db.Users().Count(ctx, r.db)
}

// GetByID returns a user by userID
func (r *authRepo) GetByID(ctx context.Context, userID string) (*db.User, error) {
	return db.Users(qm.Where("user_id = ?", userID)).One(ctx, r.db)
}
