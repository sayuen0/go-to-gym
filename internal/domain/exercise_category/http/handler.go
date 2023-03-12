package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise_category"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
	"net/http"
)

type exerciseCategoryHandlers struct {
	cfg *config.Config
	lg  logger.Logger
	uc  exercise_category.UseCase
}

func NewExerciseCategoryHandlers(
	cfg *config.Config,
	lg logger.Logger,
	uc exercise_category.UseCase,
) exercise_category.Handlers {
	return &exerciseCategoryHandlers{
		cfg: cfg,
		lg:  lg,
		uc:  uc,
	}
}

// Create godoc
// @Summary Create new exercise category
// @Description create new exercise category, return exercise category
// @Tags ExerciseCategory
// @Produces json
// @Accept json
// @Success 201 {object} models.ExerciseCategory
// @Router /auth/register [post]
func (h *exerciseCategoryHandlers) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		cat := &models.ExerciseCategoryCreateRequest{}
		if err := utils.ReadRequest(c, cat); err != nil {
			c.JSON(http.StatusBadRequest, httperrors.BadRequest(err))
			return
		}

		user, err := utils.GetLoginUser(c)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}
		cat.UserUUID = user.UserID

		createdCategory, err := h.uc.Create(ctx, cat)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusCreated, createdCategory)
	}
}

// List godoc
// @Summary List exercise categories
// @Description list exercise categories, return exercise category
// @Tags ExerciseCategory
// @Produces json
// @Accept json
// @Success 200 {object} models.ExerciseCategoryList
// @Router /auth/register [post]
func (h *exerciseCategoryHandlers) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user, err := utils.GetLoginUser(c)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		categories, err := h.uc.List(ctx, user.UserID)
		if err != nil {
			utils.LogResponseError(c, h.lg, err)
			c.JSON(httperrors.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, categories)
	}
}
