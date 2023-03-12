package models

import "github.com/sayuen0/go-to-gym/internal/models/db"

type ExerciseCategoryCreateRequest struct {
	UserUUID    string `json:"-"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

type ExerciseCategoryUpdateRequest struct {
	ID          int    `json:"id" validate:"required"`
	UserUUID    string `json:"-"`
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
	ExerciseCategories []*ExerciseCategory `json:"exerciseCategories"`
}

func NEwExerciseCategoryList(entities []*db.ExerciseCategory) *ExerciseCategoryList {
	list := &ExerciseCategoryList{}
	categories := make([]*ExerciseCategory, 0, len(entities))
	for _, e := range entities {
		categories = append(categories, NewExerciseCategory(e))
	}
	list.ExerciseCategories = categories
	return list
}
