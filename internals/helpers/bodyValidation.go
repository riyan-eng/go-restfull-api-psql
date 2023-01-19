package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/riyan-eng/go-restfull-api-psql/internals/models"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateNote(note models.Note) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(note)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
