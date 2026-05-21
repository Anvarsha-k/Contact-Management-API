package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

//validates request payloads

func ValidateStruct(payload interface{}) error {
	return validate.Struct(payload)
}