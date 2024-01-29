package utilities

import (
	
	// third party package
	"github.com/go-playground/validator/v10"
)

// function for the struct validation based on the struct tags
func ValidateStruct(data any) error {
	var validate = validator.New()
	return validate.Struct(data)
}
