package exercise

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
)

type Repository interface {
	Create(ctx context.Context, req *models.ExerciseCreateRequest) (*models.Exercise, error)
}
