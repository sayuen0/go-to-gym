package models

import "github.com/sayuen0/go-to-gym/internal/models/db"

type ExerciseCategoryCreateRequest struct {
	UUID        string `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

type ExerciseCategoryUpdateRequest struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

type ExerciseCategory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewExerciseCategory(e *db.ExerciseCategory) *ExerciseCategory {
	return &ExerciseCategory{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description.String,
	}
}

type ExerciseCategoryList struct {
	Paging             `json:"paging"`
	ExerciseCategories []*ExerciseCategory `json:"trainingCategories"`
}
