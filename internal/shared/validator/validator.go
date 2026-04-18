package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	validatorv10 "github.com/go-playground/validator/v10"
)

var validate = validatorv10.New()

func ValidateStruct(payload any) error {
	if err := validate.Struct(payload); err != nil {
		validationErrors, ok := err.(validatorv10.ValidationErrors)
		if !ok {
			return kernel.NewValidationError("invalid request payload")
		}

		messages := make([]string, 0, len(validationErrors))
		for _, fieldErr := range validationErrors {
			fieldName := fieldErr.Field()
			if parsed := jsonFieldName(payload, fieldErr.StructField()); parsed != "" {
				fieldName = parsed
			}
			messages = append(messages, fmt.Sprintf("%s failed on '%s'", fieldName, fieldErr.Tag()))
		}
		return kernel.NewValidationError(strings.Join(messages, ", "))
	}

	return nil
}

func jsonFieldName(payload any, structField string) string {
	t := reflect.TypeOf(payload)
	if t == nil {
		return ""
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}

	field, ok := t.FieldByName(structField)
	if !ok {
		return ""
	}
	tag := field.Tag.Get("json")
	if tag == "" {
		return ""
	}
	parts := strings.Split(tag, ",")
	if len(parts) == 0 || parts[0] == "" || parts[0] == "-" {
		return ""
	}
	return parts[0]
}
