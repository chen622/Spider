package utils

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	Validate = NewValidate()
)

func NewValidate() *validator.Validate {
	return validator.New()
}
