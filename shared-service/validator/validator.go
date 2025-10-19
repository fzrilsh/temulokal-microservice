package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ValidationErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

// Init initializes the validator instance
func Init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// ValidateStruct validates struct using go-playground/validator
func ValidateStruct(data interface{}) []ValidationErrorResponse {
	if validate == nil {
		Init()
	}

	err := validate.Struct(data)
	if err != nil {
		var errors []ValidationErrorResponse
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationErrorResponse{
				Field: e.Field(),
				Tag:   e.Tag(),
			})
		}

		return errors
	}

	return nil
}

// GetValidator returns the validator instance (useful for custom rules)
func GetValidator() *validator.Validate {
	if validate == nil {
		Init()
	}
	return validate
}
