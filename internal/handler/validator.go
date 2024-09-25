package handler

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func validateStruct(s interface{}) error {
	return validate.Struct(s)
}
