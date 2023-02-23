package session

import (
	"context"

	"github.com/sayuen0/go-to-gym/internal/models"
)

// UseCase defines the interface for session management use cases.
type UseCase interface {
	// CreateSession creates a new session with the provided data and expiration time.
	CreateSession(ctx context.Context, sess *models.Session, expires int) (string, error)
	// DeleteByID deletes a session by its ID.
	DeleteByID(ctx context.Context, id string) error
	// GetSessionByID retrieves a session by its ID.
	GetSessionByID(ctx context.Context, id string) (*models.Session, error)
}
