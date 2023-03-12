package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
	"net/http"
)

type exerciseHandlers struct {
	cfg *config.Config
	lg  logger.Logger
	uc  exercise.UseCase
}

func NewExerciseHandlers(
	cfg *config.Config,
	lg logger.Logger,
	uc exercise.UseCase,
) exercise.Handlers {
	return &exerciseHandlers{cfg: cfg, lg: lg, uc: uc}
}

func (h *exerciseHandlers) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		req := &models.ExerciseCreateRequest{}
		if err := utils.ReadRequest(c, req); err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		user, err := utils.GetLoginUser(c)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}
		req.UserUUID = user.UserID

		createdExercise, err := h.uc.Create(ctx, req)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusCreated, createdExercise)
	}
}

func (h *exerciseHandlers) List() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *exerciseHandlers) Get() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *exerciseHandlers) Update() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (h *exerciseHandlers) Delete() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
