package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/http_errors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
	"net/http"
)

type authHandlers struct {
	cfg    *config.Config
	uc     auth.UseCase
	sessUC session.UseCase
	lg     logger.Logger
}

func NewAuthHandlers(
	cfg *config.Config,
	uc auth.UseCase,
	sessUc session.UseCase,
	lg logger.Logger,
) auth.Handlers {
	return &authHandlers{cfg: cfg, uc: uc, sessUC: sessUc, lg: lg}
}

// Register godoc
// @Summary Register new user
// @Description register new user, return user and token
// @Tags Auth
// @Accept json
// @Produces json
// Success 201 {object} models.User
// @Route /auth/register [post]
func (h *authHandlers) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			c.JSON(http_errors.ErrorResponse(err))
			return
		}

		createdUser, err := h.uc.Register(ctx, user)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(http_errors.ErrorResponse(err))
			return
		}

		sess, err := h.sessUC.
			c.JSON(http.StatusCreated, createdUser)
		return
	}
}

func (h *authHandlers) Login() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *authHandlers) Logout() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *authHandlers) GetUsers() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *authHandlers) GetUserByID() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *authHandlers) GetMe() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *authHandlers) Update() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *authHandlers) Delete() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
