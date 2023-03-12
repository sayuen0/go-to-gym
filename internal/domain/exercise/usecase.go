package exercise

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

// UseCase provides the business logic for managing exercise.
type UseCase interface {
	// Create creates a new exercise with the given request.
	Create(ctx context.Context, req *models.ExerciseCreateRequest) (*models.Exercise, error)

	// List fetches all exercise that match the given request and pagination parameters.
	List(ctx context.Context, req *utils.PaginationRequest) (*models.ExercisesList, error)

	// Get retrieves the exercise with the given ID.
	Get(ctx context.Context, id int64) (*models.Exercise, error)

	// Update updates the exercise with the given request.
	Update(ctx context.Context, req *models.ExerciseUpdateRequest) (*models.Exercise, error)

	// Delete deletes the exercise with the given ID.
	Delete(ctx context.Context, id int64) error
}
