package usecase

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/http_errors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

type authUC struct {
	cfg       *config.Config
	lg        logger.Logger
	authRepo  auth.Repository
	redisRepo auth.RedisRepository
}

func NewAuthUseCase(cfg *config.Config, lg logger.Logger, authRepo auth.Repository, redisRepo auth.RedisRepository) auth.UseCase {
	return &authUC{cfg: cfg, lg: lg, authRepo: authRepo, redisRepo: redisRepo}
}

func (a authUC) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	existsUser, err := u.authRepo.FindByEmail(ctx, user)
	if existsUser != nil && err == nil {
		return nil, http_errors.BadRequest(err)
	}

	if err := user.PrepareCreate(); err != nil {
		return nil, http_errors.BadRequest(errors.Wrap(err, "authUC.Register.PrepareCreate"))
	}

	createdUser, err := u.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	// TODO: 個々のUseCaseではやりたくない。レスポンス加工ミドルか何かで行いたい
	createdUser.SanitizePassword()

	token, err := utils.GenerateJWTToken(createdUser, u.cfg)
	if err != nil {
		return nil, http_errors.InternalServerError(errors.Wrap(err, "authUC.Register.GenerateJWTToken"))
	}
	return &models.UserWithToken{
		User:  createdUser,
		Token: token,
	}, nil

	return nil, nil
}
