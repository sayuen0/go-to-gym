package repository

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise_category"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repository struct {
	cfg *config.Config
	db  *sql.DB
}

func NewExerciseCategoryRepo(
	cfg *config.Config,
	db *sql.DB,
) exercise_category.Repository {
	return &repository{cfg: cfg, db: db}
}

func (rp *repository) Create(ctx context.Context, req *models.ExerciseCategoryCreateRequest, userID int,
) (*db.ExerciseCategory, error) {

	c := &db.ExerciseCategory{
		UserID: userID,
		Name:   req.Name,
		Description: func(s string) null.String {
			if s == "" {
				return null.String{"", false}
			}
			return null.String{s, true}
		}(req.Description),
	}
	if err := c.Insert(ctx, rp.db, boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "insert exercise category error")
	}

	return c, nil
}

func (rp *repository) Get(ctx context.Context, id int64) (*db.ExerciseCategory, error) {
	//TODO implement me
	panic("implement me")
}

// GetByUserAndName return true if an exercise category exists with user-id and name
func (rp *repository) GetByUserAndName(ctx context.Context, userID int, name string) (*db.ExerciseCategory, error) {
	return db.ExerciseCategories(qm.Where("user_id = ? AND name = ?", userID, name)).One(ctx, rp.db)
}
