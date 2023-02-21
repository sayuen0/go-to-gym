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

func NewAuthRepo(db *sql.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, user *models.UserCreateRequest) (*models.User, error) {
	u := &db.User{
		UUID:           utils.NewUUIDStr(),
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}

	if err := u.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "authRepo.Register.Insert")
	}

	return models.NewUser(u), nil
}

func (r *authRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	u, err := db.Users(qm.Where("email = ? ", email)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return models.NewUser(u), nil
}
