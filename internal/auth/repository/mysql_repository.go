package repository

import (
	"context"
	"database/sql"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
)

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) auth.Repository {
	return &authRepo{db: db}
}

func (a authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	u := &db.User{}
	if err :=
}
