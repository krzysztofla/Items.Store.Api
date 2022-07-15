package validation

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// validation rules here
var uniqueid validator.Func = func(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(uuid.UUID)
	if ok {
		_, err := uuid.Parse(data.String())
		if err == nil {
			return false
		}
	}

	return true
}
