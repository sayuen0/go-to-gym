package utils

import (
	"context"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct returns errors if s have invalid fields
func ValidateStruct(ctx context.Context, s any) error {
	return validate.StructCtx(ctx, s)
}
