package usecase

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise_category"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

type exerciseUC struct {
	cfg    *config.Config
	lg     logger.Logger
	repo   exercise.Repository
	tcRepo exercise_category.Repository
}

func NewExerciseUseCase(
	cfg *config.Config,
	lg logger.Logger,
	repo exercise.Repository,
	tcRepo exercise_category.Repository,
) exercise.UseCase {
	return &exerciseUC{
		cfg: cfg, lg: lg, repo: repo, tcRepo: tcRepo,
	}
}

// Create creates a new
func (e *exerciseUC) Create(ctx context.Context, req *models.ExerciseCreateRequest) (*models.Exercise, error) {
	if err := e.PrepareForCreate(ctx, req); err != nil {
		e.lg.Error("prepare for create", logger.Error(err), logger.Struct("req", req))
		return nil, httperrors.BadRequest(err)
	}

	// 作成
	createdExercise, err := e.repo.Create(ctx, req)
	if err != nil {
		e.lg.Error("create", logger.Error(err), logger.Struct("req", req))
		return nil, errors.Wrap(err, "createA")
	}

	return createdExercise, nil
}

// PrepareForCreate returns error if request can be processed
func (e *exerciseUC) PrepareForCreate(ctx context.Context, req *models.ExerciseCreateRequest) error {
	// カテゴリが存在するかを確認
	_, err := e.tcRepo.Get(ctx, req.CategoryID)
	if err != nil {
		return errors.Wrap(err, "tc repo get")
	}

	return nil
}

func (e *exerciseUC) List(ctx context.Context, req *utils.PaginationRequest) (*models.ExercisesList, error) {
	//TODO implement me
	panic("implement me")
}

func (e *exerciseUC) Get(ctx context.Context, id int64) (*models.Exercise, error) {
	//TODO implement me
	panic("implement me")
}

func (e *exerciseUC) Update(ctx context.Context, req *models.ExerciseUpdateRequest) (*models.Exercise, error) {
	//TODO implement me
	panic("implement me")
}

func (e *exerciseUC) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
