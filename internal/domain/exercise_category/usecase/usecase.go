package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise_category"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
)

type categoryUC struct {
	cfg      *config.Config
	lg       logger.Logger
	repo     exercise_category.Repository
	authRepo auth.Repository
}

// NewExerciseCategoryUseCase creates a new exercise_category.UseCase
func NewExerciseCategoryUseCase(
	cfg *config.Config,
	lg logger.Logger,
	repo exercise_category.Repository,
	authRepo auth.Repository,
) exercise_category.UseCase {
	return &categoryUC{
		cfg: cfg, lg: lg, repo: repo, authRepo: authRepo,
	}
}

// Create creates a new exercise category
func (uc *categoryUC) Create(ctx context.Context, req *models.ExerciseCategoryCreateRequest) (*models.ExerciseCategory, error) {
	if err := uc.PrepareForCreate(ctx, req); err != nil {
		uc.lg.Error("prepare for create ", logger.Error(err), logger.Struct("req", req))
		return nil, httperrors.BadRequest(err)
	}

	user, err := uc.authRepo.GetByUUID(ctx, req.UserUUID)
	if err != nil {
		uc.lg.Error("get user by id", logger.Error(err), logger.Struct("req", req))
		return nil, httperrors.InternalServerError(err)
	}

	createdCategory, err := uc.repo.Create(ctx, req, user.ID)
	if err != nil {
		uc.lg.Error("create exercise category", logger.Error(err), logger.Struct("req", req))
		return nil, err
	}

	return models.NewExerciseCategory(createdCategory), nil
}

// PrepareForCreate returns error if request is executable
func (uc *categoryUC) PrepareForCreate(ctx context.Context, req *models.ExerciseCategoryCreateRequest) error {
	// ユーザの存在チェック
	user, err := uc.authRepo.GetByUUID(ctx, req.UserUUID)
	if err != nil {
		return errors.Wrap(err, "user not found")
	}

	// 同一ユーザで、同一名称を持っていないかチェック
	exerciseCategory, err := uc.repo.GetByUserAndName(ctx, user.ID, req.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 不在以外のエラー
		return errors.Wrap(err, "exists user-id and name")
	} else if exerciseCategory != nil {
		// 存在した
		return fmt.Errorf("user has same named category")
	}

	return nil
}

func (uc *categoryUC) List(ctx context.Context, userID string) (*models.ExerciseCategoryList, error) {
	// ユーザの存在チェック
	user, err := uc.authRepo.GetByUUID(ctx, userID)
	if err != nil {
		msg := "user not found"
		uc.lg.Error("", logger.Error(err), logger.String("user_id", userID))
		return nil, errors.Wrap(err, msg)
	}

	// ユーザ配下のカテゴリ取得
	categories, err := uc.repo.GetByUserID(ctx, user.ID)
	if err != nil {
		msg := "get exercise categories by user id"
		uc.lg.Error(msg, logger.Error(err), logger.String("uuid", userID))
		return nil, errors.Wrap(err, msg)
	}

	return models.NEwExerciseCategoryList(categories), nil
}
