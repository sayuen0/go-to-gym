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
	cfg      *config.Config
	lg       logger.Logger
	authRepo auth.Repository
}

func NewAuthUseCase(cfg *config.Config, lg logger.Logger, authRepo auth.Repository) auth.UseCase {
	return &authUC{cfg: cfg, lg: lg, authRepo: authRepo}
}

func (u *authUC) Register(ctx context.Context, req *models.UserCreateRequest) (*models.UserWithToken, error) {
	existsUser, err := u.authRepo.FindByEmail(ctx, req.Email)
	if existsUser != nil && err == nil {
		return nil, http_errors.BadRequest(errors.New("already registered email"))
	}

	if err := req.PrepareForCreate(); err != nil {
		return nil, http_errors.BadRequest(errors.Wrap(err, "authUC.Register.PrepareCreate"))
	}

	createdUser, err := u.authRepo.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWTToken(createdUser.Email, createdUser.UUID, u.cfg)
	if err != nil {
		return nil, http_errors.InternalServerError(errors.Wrap(err, "authUC.Register.GenerateJWTToken"))
	}
	return &models.UserWithToken{
		User:  models.NewUser(createdUser),
		Token: token,
	}, nil
}

// Login user, returns user model with jwt token
func (u *authUC) Login(ctx context.Context, req *models.UserLoginRequest) (*models.UserWithToken, error) {
	dbUser, err := u.authRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := models.CompareUserPassword(req.Password, dbUser.HashedPassword); err != nil {
		return nil, http_errors.Unauthorized(errors.Wrap(err, "authUC.Login.CompareUserPassword"))
	}

	token, err := utils.GenerateJWTToken(dbUser.Email, dbUser.UUID, u.cfg)
	if err != nil {
		return nil, http_errors.InternalServerError(errors.Wrap(err, "authUC.GetUsers.GenerateJWTToken"))
	}

	return &models.UserWithToken{
		User:  models.NewUser(dbUser),
		Token: token,
	}, nil
}
