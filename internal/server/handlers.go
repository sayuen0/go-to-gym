package server

import (
	"github.com/gin-gonic/gin"

	authHttp "github.com/sayuen0/go-to-gym/internal/auth/http"
	authRepo "github.com/sayuen0/go-to-gym/internal/auth/repository"
	authUseCase "github.com/sayuen0/go-to-gym/internal/auth/usecase"
	sessRepo "github.com/sayuen0/go-to-gym/internal/session/repository"
	sessUseCase "github.com/sayuen0/go-to-gym/internal/session/usecase"
	"net/http"
)

func (s *server) Handle(r *gin.Engine) error {
	// TODO: init repositories
	authRp := authRepo.NewAuthRepo(s.cfg, s.db)
	authRedisRp := authRepo.NewRedisRepo(s.redisClient)
	sessRp := sessRepo.NewSessionRepo(s.redisClient, s.cfg)

	// TODO init use cases
	authUC := authUseCase.NewAuthUseCase(s.cfg, s.lg, authRp, authRedisRp)
	sessUC := sessUseCase.NewSessionUseCase(s.cfg, sessRp)

	// TODO use middlewares

	// TODO init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, s.lg, authUC, sessUC)

	authGroup := r.Group("/auth")
	authHttp.MapAuthRoutes(authGroup, authHandlers)

	health := r.Group("/health")
	health.GET("/", func(c *gin.Context) {
		// TODO: identify health check ID
		s.lg.Info("Health check")
		c.String(http.StatusOK, "OK")
	})
	return nil
}
