package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise_category"
	"github.com/sayuen0/go-to-gym/internal/middleware"
)

func MapExerciseCategoryRoutes(
	rg *gin.RouterGroup, h exercise_category.Handlers, mw *middleware.Wrapper) {
	rg.GET("/", h.List())
	rg.POST("/", h.Create())
}
