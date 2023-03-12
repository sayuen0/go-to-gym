package repository

import (
	"context"
	"database/sql"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type exerciseRepo struct {
	cfg *config.Config
	db  *sql.DB
}

// NewExerciseRepo creates a new exercise repository
func NewExerciseRepo(
	cfg *config.Config,
	db *sql.DB,
) exercise.Repository {
	return &exerciseRepo{
		cfg: cfg, db: db,
	}
}

// Create creates a new exercise
func (rp *exerciseRepo) Create(
	ctx context.Context,
	req *models.ExerciseCreateRequest,
	userID int,
) (*db.Exercise, error) {
	e := &db.Exercise{
		Name: req.Name,
		Description: func(s string) null.String {
			if s == "" {
				return null.String{"", false}
			}
			return null.String{s, true}
		}(req.Description),
		UserID:     userID,
		CategoryID: req.CategoryID,
	}

	if err := e.Insert(ctx, rp.db, boil.Infer()); err != nil {
		return nil, err
	}

	return e, nil
}

// GetByCategoryAndName returns a exercised by the given categoryID and name
func (rp *exerciseRepo) GetByCategoryAndName(ctx context.Context, categoryID int, name string) (*db.Exercise, error) {
	return db.Exercises(qm.Where("category_id = ? and name = ?", categoryID, name)).One(ctx, rp.db)
}
