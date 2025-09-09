package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidate[T any](c *fiber.Ctx, payload *T) error {

	err := c.BodyParser(payload)
	if err != nil {
		return errors.New("invalid json")
	}

	err = validate.Struct(payload)
	if err != nil {
		return errors.New(formatValidationErrors(err, payload))
	}

	return nil
}

func formatValidationErrors(rawError error, payload any) string {

	errs, ok := rawError.(validator.ValidationErrors)
	if !ok {
		return rawError.Error()
	}

	// Get the type of the payload. If it's a pointer, dereference it
	t := reflect.TypeOf(payload)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var msg string = ""

	for _, e := range errs {

		// Get the struct field by its name
		field, found := t.FieldByName(e.StructField())

		if !found {
			// Use the struct field name as a fallback
			msg += fmt.Sprintf("%s: %s\n", e.StructField(), e.Tag())
		} else {

			// Extract the JSON tag
			jsonTag := field.Tag.Get("json")
			jsonName := strings.Split(jsonTag, ",")[0]

			// Determine the name to use: JSON tag if valid, otherwise the struct field name
			nameToUse := field.Name
			if jsonName != "" && jsonName != "-" {
				nameToUse = jsonName
			}

			msg += fmt.Sprintf("%s: %s", nameToUse, e.Tag())
			p := e.Param()
			if p != "" {
				msg += fmt.Sprintf("= %s", p)
			}
			msg += "\n"
		}
	}

	return msg
}
