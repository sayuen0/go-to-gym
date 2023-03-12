package usecase

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise_category"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
)

type exerciseUC struct {
	cfg      *config.Config
	lg       logger.Logger
	repo     exercise.Repository
	authRepo auth.Repository
	exCtRepo exercise_category.Repository
}

// NewExerciseUseCase creates a new exercise use case
func NewExerciseUseCase(
	cfg *config.Config,
	lg logger.Logger,
	repo exercise.Repository,
	authRepo auth.Repository,
	ecRepo exercise_category.Repository,
) exercise.UseCase {
	return &exerciseUC{
		cfg: cfg, lg: lg, repo: repo, authRepo: authRepo, exCtRepo: ecRepo,
	}
}

// Create creates a new exercise
func (uc *exerciseUC) Create(ctx context.Context, req *models.ExerciseCreateRequest) (*models.Exercise, error) {
	if err := uc.PrepareForCreate(ctx, req); err != nil {
		uc.lg.Error("prepare for create", logger.Error(err), logger.Struct("req", req))
		return nil, httperrors.BadRequest(err)
	}

	user, err := uc.authRepo.GetByUUID(ctx, req.UserUUID)
	if err != nil {
		uc.lg.Error("get user", logger.Error(err), logger.Struct("req", req))
		return nil, httperrors.InternalServerError(err)
	}

	createdExercise, err := uc.repo.Create(ctx, req, user.ID)
	if err != nil {
		uc.lg.Error("create", logger.Error(err), logger.Struct("req", req))
		return nil, errors.Wrap(err, "createA")
	}

	return models.NewExercise(createdExercise), nil
}

// PrepareForCreate returns error if request can be processed
func (uc *exerciseUC) PrepareForCreate(ctx context.Context, req *models.ExerciseCreateRequest) error {
	// カテゴリが存在するかを確認
	_, err := uc.exCtRepo.Get(ctx, req.CategoryID)
	if err != nil {
		return errors.Wrap(err, "exCt repo get")
	}

	// 同一カテゴリ、同一名称の種目が存在するかを確認
	if existingExercise, err := uc.repo.GetByCategoryAndName(ctx, req.CategoryID, req.Name); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "repo get by category and name")
	} else if existingExercise != nil {
		return errors.New("exercise duplicated by category and name")
	}

	return nil
}

func (uc *exerciseUC) GetByUserID(ctx context.Context, userUUID string) (*models.ExercisesList, error) {
	user, err := uc.authRepo.GetByUUID(ctx, userUUID)
	if err != nil {
		uc.lg.Error("get user", logger.Error(err), logger.String("user_id", user.UserID))
		return nil, httperrors.InternalServerError(err)
	}

	exercises, err := uc.repo.GetByUserID(ctx, user.ID)
	if err != nil {
		uc.lg.Error("get exercises", logger.Error(err), logger.String("user_id", user.UserID))
		return nil, httperrors.InternalServerError(err)
	}

	return models.NewExercisesList(exercises), nil
}

func (uc *exerciseUC) Get(ctx context.Context, id int64) (*models.Exercise, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *exerciseUC) Update(ctx context.Context, req *models.ExerciseUpdateRequest) (*models.Exercise, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *exerciseUC) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
