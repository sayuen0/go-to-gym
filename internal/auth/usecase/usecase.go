package usecase

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

const (
	cacheDuration = 3600 * 3
)

type authUC struct {
	cfg       *config.Config
	lg        logger.Logger
	authRepo  auth.Repository
	redisRepo auth.RedisRepository
}

// NewAuthUseCase is a constructor for authUC
func NewAuthUseCase(
	cfg *config.Config,
	lg logger.Logger,
	authRepo auth.Repository,
	redisRepo auth.RedisRepository) auth.UseCase {
	return &authUC{cfg: cfg, lg: lg, authRepo: authRepo, redisRepo: redisRepo}
}

// Register registers a new user
func (u *authUC) Register(ctx context.Context, req *models.UserCreateRequest) (*models.UserWithToken, error) {
	existsUser, err := u.authRepo.FindByEmail(ctx, req.Email)
	if existsUser != nil && err == nil {
		return nil, httperrors.BadRequest(errors.New("already registered email"))
	}

	if err := req.PrepareForCreate(); err != nil {
		return nil, httperrors.BadRequest(errors.Wrap(err, "authUC.Register.PrepareCreate"))
	}

	createdUser, err := u.authRepo.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWTToken(createdUser.Email, createdUser.UserID, u.cfg)
	if err != nil {
		return nil, httperrors.InternalServerError(errors.Wrap(err, "authUC.Register.GenerateJWTToken"))
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

	password := utils.PasswordVerifier{
		Salt:   dbUser.Salt,
		Pepper: u.cfg.Server.SecretSalt,
	}
	if err := password.ComparePassword(req.Password, dbUser.HashedPassword); err != nil {
		return nil, httperrors.Unauthorized(errors.Wrap(err, "authUC.Login.ComparePassword"))
	}

	token, err := utils.GenerateJWTToken(dbUser.Email, dbUser.UserID, u.cfg)
	if err != nil {
		return nil, httperrors.InternalServerError(errors.Wrap(err, "authUC.GetUsers.GenerateJWTToken"))
	}

	return &models.UserWithToken{
		User:  models.NewUser(dbUser),
		Token: token,
	}, nil
}

// GetUsers returns *models.UsersList queried with *utils.PaginationRequest
func (u *authUC) GetUsers(ctx context.Context, req *utils.PaginationRequest) (*models.UsersList, error) {
	totalCount, err := u.authRepo.GetCount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "authRepo.GetUsers.GetCount")
	}
	users, err := u.authRepo.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	return models.NewUsersList(users, models.Paging{
		Page:  req.Page,
		Size:  req.Size,
		Total: totalCount,
	}), nil
}

// GetByID returns a user found by its UUID
func (u *authUC) GetByID(ctx context.Context, userID string) (*models.User, error) {
	cachedUser, err := u.redisRepo.GetByID(ctx, userID)
	if err != nil {
		u.lg.Error("authUC.GetByID.redisRepo.GetByID", logger.Error(err))
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	user, err := u.authRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userModel := models.NewUser(user)
	if err := u.redisRepo.SetUser(ctx, userID, cacheDuration, userModel); err != nil {
		u.lg.Error("authUC.GetByID.SetUser", logger.Error(err))
	}

	return userModel, nil
}
