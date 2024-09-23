package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateSignInRequest(input interface{}) error {
	return validate.Struct(input)
}
