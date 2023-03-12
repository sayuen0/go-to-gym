package exercise_category

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, req *models.ExerciseCategoryCreateRequest) (*models.ExerciseCategory, error)
	List(ctx context.Context, userID string) (*models.ExerciseCategoryList, error)
}
