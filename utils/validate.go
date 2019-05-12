package utils

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	Validate = New()
)

func New() *validator.Validate {
	return validator.New()
}
