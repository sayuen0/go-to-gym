package middleware

import (
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/session"
)

// Wrapper is a group of middleware
type Wrapper struct {
	cfg    *config.Config
	lg     logger.Logger
	sessUC session.UseCase
	authUC auth.UseCase
}

// NewMiddlewareWrapper creates a new middleware wrapper
func NewMiddlewareWrapper(
	cfg *config.Config,
	lg logger.Logger,
	sessUC session.UseCase,
	authUC auth.UseCase,
) *Wrapper {
	return &Wrapper{cfg: cfg, lg: lg, sessUC: sessUC, authUC: authUC}
}
