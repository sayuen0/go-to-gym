package models

// ---------------------------------------------------------------------------------------------------------------------
// exercise create request

type ExerciseCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	CategoryID  int64  `json:"category_id" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

// ---------------------------------------------------------------------------------------------------------------------
// exercise update request

type ExerciseUpdateRequest struct {
	ID          int64  `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	CategoryID  int64  `json:"category_id" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

// ---------------------------------------------------------------------------------------------------------------------
// exercise response

type Exercise struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CategoryID  int64  `json:"category_id"`
	Description string `json:"description"`
}

// ---------------------------------------------------------------------------------------------------------------------
// exercise list response

type ExercisesList struct {
	Paging    `json:"paging"`
	Exercises []*Exercise `json:"exercise"`
}
