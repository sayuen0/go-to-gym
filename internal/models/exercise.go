package models

import "github.com/sayuen0/go-to-gym/internal/models/db"

// ---------------------------------------------------------------------------------------------------------------------
// exercise create request

type ExerciseCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	UserUUID    string `json:"-"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

// ---------------------------------------------------------------------------------------------------------------------
// exercise update request

type ExerciseUpdateRequest struct {
	ID          int    `json:"id" validate:"required"`
	UserUUID    string `json:"-"`
	Name        string `json:"name" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

// ---------------------------------------------------------------------------------------------------------------------
// exercise response

type Exercise struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
}

func NewExercise(e *db.Exercise) *Exercise {
	return &Exercise{
		ID:          e.ID,
		Name:        e.Name,
		CategoryID:  e.CategoryID,
		Description: e.Description.String,
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// exercise list response

type ExercisesList struct {
	Exercises []*Exercise `json:"exercise"`
}

// NewExercisesList return a list of Exercise objects
func NewExercisesList(entities []*db.Exercise) *ExercisesList {
	es := make([]*Exercise, 0)
	for _, e := range entities {
		es = append(es, NewExercise(e))
	}
	return &ExercisesList{
		Exercises: es,
	}

}
