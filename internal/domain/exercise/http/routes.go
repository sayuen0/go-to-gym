package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise"
	"github.com/sayuen0/go-to-gym/internal/middleware"
)

func MapExerciseRoutes(rg *gin.RouterGroup, h exercise.Handlers, mw *middleware.Wrapper) {
	rg.POST("/", h.Create())
	rg.GET("/", h.List())

}
