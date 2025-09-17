package middleware

import (
	"strings"

	"github.com/HivenOrg/Hiven/config"
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB, cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) < 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token := tokenParts[1]
		userId, err := utils.ValidateJWT(token, cfg.JWT_SECRET_KEY)

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Checking if user exists in database
		user, err := utils.GetUserById(userId, db)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "invalid token"})
		}

		// Fetching Hive IDs of this user
		hiveIDs, err := utils.GetHiveIDsForUser(userId, db)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Locals("user_id", user.ID)
		c.Locals("hive_ids", hiveIDs)

		return c.Next()
	}
}
