package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/session"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

// authHandlers auth handlers
type authHandlers struct {
	cfg    *config.Config
	uc     auth.UseCase
	sessUC session.UseCase
	lg     logger.Logger
}

// NewAuthHandlers is a constructor of authHandlers
func NewAuthHandlers(
	cfg *config.Config,
	lg logger.Logger,
	uc auth.UseCase,
	sessUc session.UseCase,
) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		lg:     lg,
		uc:     uc,
		sessUC: sessUc,
	}
}

// Register godoc
// @Summary Register new user
// @Description register new user, return user and token
// @Tags Auth
// @Produces json
// @Accept json
// @Success 201 {object} models.UserWithToken
// @Router /auth/register [post]
func (h *authHandlers) Register() gin.HandlerFunc {
	// TODO: add failure situation on swagger
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := &models.UserCreateRequest{}
		if err := utils.ReadRequest(c, user); err != nil {
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		createdUser, err := h.uc.Register(ctx, user)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		sess, err := h.sessUC.CreateSession(ctx,
			&models.Session{UserID: createdUser.User.UserID},
			h.cfg.Session.Expire)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}
		utils.CreateSessionCookie(c, h.cfg, sess)

		c.JSON(http.StatusCreated, createdUser)
	}
}

// Login godoc
// @Summary Login as a user
// @Description login user, returns user and set session
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.UserWithToken
// @Router /auth/login [post]
func (h *authHandlers) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		login := &models.UserLoginRequest{}
		if err := utils.ReadRequest(c, login); err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		userWithToken, err := h.uc.Login(ctx, &models.UserLoginRequest{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		sess, err := h.sessUC.CreateSession(ctx,
			&models.Session{UserID: userWithToken.User.UserID},
			h.cfg.Session.Expire)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}
		utils.CreateSessionCookie(c, h.cfg, sess)

		c.JSON(http.StatusCreated, userWithToken)
	}
}

// Logout godoc
// @Summary Logout user
// @Description logout user removing session
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 500 {string} internal server error
// @Router /auth/logout [post]
func (h *authHandlers) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		sessionID, err := c.Cookie(h.cfg.Session.Name)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				c.JSON(http.StatusUnauthorized, httperrors.Unauthorized(err))
				return
			}
			utils.LogResponseError(c, h.lg, err)
			c.JSON(http.StatusInternalServerError, httperrors.InternalServerError(err))
			return
		}

		if err := h.sessUC.DeleteByID(ctx, sessionID); err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		utils.DeleteSessionCookie(c, h.cfg)
		c.Status(http.StatusOK)
	}
}

// GetUsers godoc
// @Summary Get users
// @Description Get the list of all users
// @Tags Auth
// @Accept json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Produce json
// @Success 200 {object} models.UsersList
// @Failure 500 {object} httperrors.RestErr
// @Router /auth/all [get]
func (h *authHandlers) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)
		paginationReq, err := utils.GetPaginationRequest(c)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		users, err := h.uc.GetUsers(ctx, paginationReq)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetUserByID godoc
// @Summary get user by id
// @Description get string by ID
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param id path int true "user_id"
// @Success 200 {object} models.User
// @Failure 500 {object} httperrors.RestErr
// @Router /auth/{id} [get]
func (h *authHandlers) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)
		uuid := c.Param("user_id")

		user, err := h.uc.GetByID(ctx, uuid)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// GetMe godoc
// @Summary Get user by id
// @Description Get current user by id
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 500 {object} httperrors.RestErr
// @Router /auth/me [get]
func (h *authHandlers) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, found := c.Get("user")
		if !found {
			utils.LogResponseError(c, h.lg, httperrors.Unauthorized(httperrors.ErrUnauthorized))
			utils.ErrorResponseWithLog(c, h.lg, httperrors.Unauthorized(httperrors.ErrUnauthorized))
			return
		}
		user, ok := u.(*models.User)
		if !ok {
			utils.LogResponseError(c, h.lg, httperrors.InternalServerError(httperrors.ErrInternalServerError))
			utils.ErrorResponseWithLog(c, h.lg, httperrors.InternalServerError(httperrors.ErrInternalServerError))
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (h *authHandlers) Update() gin.HandlerFunc {
	// TODO implement me
	panic("implement me")
}

func (h *authHandlers) Delete() gin.HandlerFunc {
	// TODO implement me
	panic("implement me")
}
