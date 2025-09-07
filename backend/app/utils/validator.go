package utils

import (
	"errors"
	"fmt"

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
		return errors.New(formatValidationErrors(err))
	}

	return nil
}

func formatValidationErrors(rawError error) string {
	errors, assertionSuccess := rawError.(validator.ValidationErrors) // errors is a slice of all individual errors

	if assertionSuccess {
		msg := ""
		for _, e := range errors {
			msg += fmt.Sprintf("%s: %s", e.Field(), e.ActualTag())
			if e.Param() != "" {
				msg += fmt.Sprintf("=%s", e.Param())
			}
			msg += "\n"
		}
		return msg
	}

	return rawError.Error()
}
