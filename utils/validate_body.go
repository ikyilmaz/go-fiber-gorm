package utils

import (
	"fiber-rest-api/validators"
	"github.com/gofiber/fiber"
)

func ValidateBody(structPtr validators.IValidator, c *fiber.Ctx) error {

	if err := c.BodyParser(structPtr); err != nil {
		return err
	}

	if err := structPtr.Validate(); err != nil {
		return err
	}

	return nil
}
