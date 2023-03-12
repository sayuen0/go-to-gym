package usecase

import (
	"context"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise_category"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type categoryUC struct {
	cfg  *config.Config
	lg   logger.Logger
	repo exercise_category.Repository
}

func NewExerciseCategoryUseCase(
	cfg *config.Config,
	lg logger.Logger,
	repo exercise_category.Repository,
) exercise_category.UseCase {
	return &categoryUC{
		cfg: cfg, lg: lg, repo: repo,
	}
}

// Create creates a new exercise category
func (uc *categoryUC) Create(ctx context.Context, req *models.ExerciseCategoryCreateRequest) (*models.ExerciseCategory, error) {
	createdCategory, err := uc.repo.Create(ctx, req)
	if err != nil {
		uc.lg.Error("create exercise category", logger.Error(err), logger.Struct("req", req))
		return nil, err
	}

	return models.NewExerciseCategory(createdCategory), nil
}
