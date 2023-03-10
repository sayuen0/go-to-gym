package session

//go:generate moq -out ./mock/${GOFILE} -pkg mock . Repository

import (
	"context"

	"github.com/sayuen0/go-to-gym/internal/models"
)

// Repository defines the methods that a session repository must implement.
type Repository interface {
	// CreateSession creates a new session for the provided user and returns the session ID.
	// The `expires` parameter specifies the session expiration time in seconds.
	Create(ctx context.Context, sess *models.Session, expires int) (string, error)

	// DeleteByID deletes the session with the provided ID.
	DeleteByID(ctx context.Context, userID string) error

	// GetSessionByID returns the session with the provided ID.
	GetByID(ctx context.Context, id string) (*models.Session, error)
}
