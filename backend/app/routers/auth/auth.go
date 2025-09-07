package auth

import (
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB) {

	router.Post("/register", func(c *fiber.Ctx) error {

		var requestBody RegistrationRequestSchema

		// Parsing and Validating
		err := utils.ParseAndValidate(c, &requestBody)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.JSON(fiber.Map{"msg": "success"})
	})
}
