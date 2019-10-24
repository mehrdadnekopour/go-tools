package mypes

import (
	"gopkg.in/go-playground/validator.v9"
)

// Validate ...
var Validate *validator.Validate

// NewValidate ...
func NewValidate() {
	Validate = validator.New()
}
