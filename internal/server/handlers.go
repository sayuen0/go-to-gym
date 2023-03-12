package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authHttp "github.com/sayuen0/go-to-gym/internal/auth/http"
	authRepo "github.com/sayuen0/go-to-gym/internal/auth/repository"
	authUseCase "github.com/sayuen0/go-to-gym/internal/auth/usecase"
	exCtHttp "github.com/sayuen0/go-to-gym/internal/domain/exercise_category/http"
	exCtRepo "github.com/sayuen0/go-to-gym/internal/domain/exercise_category/repository"
	exCtUseCase "github.com/sayuen0/go-to-gym/internal/domain/exercise_category/usecase"
	"github.com/sayuen0/go-to-gym/internal/middleware"
	sessRepo "github.com/sayuen0/go-to-gym/internal/session/repository"
	sessUseCase "github.com/sayuen0/go-to-gym/internal/session/usecase"
)

func (s *server) Handle(r *gin.Engine) error {
	// -----------------------------------------------------------------------------------------------------------------
	// repositories
	authRp := authRepo.NewAuthRepo(s.cfg, s.db)
	authRedisRp := authRepo.NewRedisRepo(s.redisClient)
	sessRp := sessRepo.NewSessionRepo(s.redisClient, s.cfg)

	exCtRp := exCtRepo.NewExerciseCategoryRepo(s.cfg, s.db)

	// -----------------------------------------------------------------------------------------------------------------
	// use-cases
	authUC := authUseCase.NewAuthUseCase(s.cfg, s.lg, authRp, authRedisRp)
	sessUC := sessUseCase.NewSessionUseCase(s.cfg, sessRp)
	exCtUC := exCtUseCase.NewExerciseCategoryUseCase(s.cfg, s.lg, exCtRp, authRp)

	// -----------------------------------------------------------------------------------------------------------------
	// middlewares
	mw := middleware.NewMiddlewareWrapper(s.cfg, s.lg, sessUC, authUC)

	r.Use(gin.Recovery())

	if s.cfg.Server.TLS {
		r.Use(mw.RedirectHTTPToHTTPS())
	}

	if s.cfg.Server.Debug {
		r.Use(mw.DebugMiddleware())
	}

	// -----------------------------------------------------------------------------------------------------------------
	// handlers
	// auth
	authHandlers := authHttp.NewAuthHandlers(s.cfg, s.lg, authUC, sessUC)
	authGroup := r.Group("/auth")
	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)

	// exercise-category
	exCtHandlers := exCtHttp.NewExerciseCategoryHandlers(s.cfg, s.lg, exCtUC)
	exCtGroup := r.Group("/exercise_categories", mw.AuthSessionMiddleware())
	exCtHttp.MapExerciseCategoryRoutes(exCtGroup, exCtHandlers, mw)

	// health check
	health := r.Group("/health")
	health.GET("/", func(c *gin.Context) {
		// TODO: identify health check ID
		s.lg.Info("Health check")
		c.String(http.StatusOK, "OK")
	})

	return nil
}
