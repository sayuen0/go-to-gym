package usecase

import (
	"context"

	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/session"
)

type sessionUC struct {
	cfg *config.Config
	rp  session.Repository
}

// NewSessionUseCase returns a new instance of the session use case.
func NewSessionUseCase(cfg *config.Config, rp session.Repository) session.UseCase {
	return &sessionUC{
		cfg: cfg,
		rp:  rp,
	}
}

// CreateSession creates a new session with the given session information and expiration
func (u *sessionUC) CreateSession(ctx context.Context, sess *models.Session, expires int) (string, error) {
	return u.rp.CreateSession(ctx, sess, expires)
}

// DeleteByID deletes the session with the given ID.
func (u *sessionUC) DeleteByID(ctx context.Context, id string) error {
	return u.rp.DeleteByID(ctx, id)
}

// GetSessionByID retrieves the session with the given ID.
func (u *sessionUC) GetSessionByID(ctx context.Context, id string) (*models.Session, error) {
	return u.rp.GetSessionByID(ctx, id)
}
