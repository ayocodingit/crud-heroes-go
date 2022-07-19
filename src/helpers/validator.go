package helpers

import (
	"github.com/go-playground/validator/v10"
)

func Validate(payload interface{}) (err error) {
	validate := validator.New()

	err = validate.Struct(payload)

	return
}
