package exercise

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/models/db"
)

type Repository interface {
	Create(ctx context.Context, req *models.ExerciseCreateRequest, userID int) (*db.Exercise, error)
	GetByCategoryAndName(ctx context.Context, CategoryID int, name string) (*db.Exercise, error)
}
