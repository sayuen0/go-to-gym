package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/domain/exercise"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
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

func (e *exerciseHandlers) List() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (e *exerciseHandlers) Get() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (e *exerciseHandlers) Update() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (e *exerciseHandlers) Delete() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (e *exerciseHandlers) Create() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
