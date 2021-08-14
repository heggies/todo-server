package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
}

var v *validator.Validate

func getValidator() *validator.Validate {
	if v == nil {
		v = validator.New()
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return v
}

func ValidateStruct(data interface{}) (errors []Error, err error) {
	if reflect.ValueOf(data).Kind() != reflect.Struct {
		err = fmt.Errorf("argument is not a struct")
		return
	}

	if err = getValidator().Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, Error{
				FailedField: err.Field(),
				Tag:         err.Tag(),
			})
		}
	}

	err = nil
	return
}
