package utils

import (
	"context"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(ctx context.Context, s any) error {
	return validate.StructCtx(ctx, s)
}
