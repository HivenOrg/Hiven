package auth

import (
	"github.com/HivenOrg/Hiven/config"
	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB, cfg *config.Config) {

	router.Post("/register", func(c *fiber.Ctx) error {

		var reqBody RegistrationRequestSchema

		// Parsing and Validating
		err := utils.ParseAndValidate(c, &reqBody)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// Email exists
		exists, err := utils.UserWithEmailExists(reqBody.Email, db)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if exists {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"msg": "email is already in use"})
		}

		// Phone number exists
		exists, err = utils.UserWithPhoneNumberExists(reqBody.PhoneNumber, db)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if exists {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"msg": "phone number is already in use"})
		}

		// Hashing password
		hashedPassword, err := utils.HashPassword(reqBody.Password)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Creating a new user
		newUser := &database.User{
			Email:       reqBody.Email,
			Password:    hashedPassword,
			FirstName:   reqBody.FirstName,
			LastName:    reqBody.LastName,
			PhoneNumber: reqBody.PhoneNumber,
		}

		err = db.Create(newUser).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Generating JWT token
		token, err := utils.GenerateJWT(newUser.ID, cfg.JWT_SECRET_KEY, cfg.TOKEN_EXPIRES_IN_HOURS)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"bearer_token": token})
	})
}
