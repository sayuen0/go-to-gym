package session

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type UseCase interface {
	CreateSession(ctx context.Context, sess *models.Session, expires int) (string, error)
}
