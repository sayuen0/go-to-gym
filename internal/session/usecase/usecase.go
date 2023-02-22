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

func NewSessionUseCase(cfg *config.Config, rp session.Repository) session.UseCase {
	return &sessionUC{
		cfg: cfg,
		rp:  rp,
	}
}

func (u *sessionUC) CreateSession(ctx context.Context, sess *models.Session, expires int) (string, error) {
	return u.rp.CreateSession(ctx, sess, expires)
}

func (u *sessionUC) DeleteByID(ctx context.Context, id string) error {
	return u.rp.DeleteByID(ctx, id)
}
