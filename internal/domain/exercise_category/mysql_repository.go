package exercise_category

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
)

type Repository interface {
	Create(ctx context.Context, req *models.ExerciseCategoryCreateRequest) (*db.ExerciseCategory, error)
	Get(ctx context.Context, id int64) (*models.ExerciseCategory, error)
	//List(ctx context.Context, req *utils.PaginationRequest) (*models.ExerciseCategoryList, error)
	//Update(ctx context.Context, req *models.ExerciseCategoryUpdateRequest) (*models.ExerciseCategory, error)
	//Delete(ctx context.Context, id int64) error
}
