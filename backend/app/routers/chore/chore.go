package chore

import (
	"slices"
	"strconv"

	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChoreRouter(router fiber.Router, db *gorm.DB) {

	router.Post("/create", func(c *fiber.Ctx) error {

		var reqBody createChore

		// Parsing and Validating
		err := utils.ParseAndValidate(c, &reqBody)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// Fetching current user details
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		hiveIDs, ok := c.Locals("hive_ids").([]uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Checking if current user is authorized to perform this action
		authorized := slices.Contains(hiveIDs, reqBody.HiveID)
		if !authorized {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Adding chore to database
		newChore := database.Chore{
			HiveID:    reqBody.HiveID,
			CreatorID: userID,
			Chore:     reqBody.Chore,
		}

		err = db.Create(&newChore).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": "Created"})
	})

	router.Get("/list", func(c *fiber.Ctx) error {

		// Parsing and Validating
		idParam := c.Query("hive_id")

		parsed, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid hive_id")
		}

		targetHiveID := uint(parsed)

		// Fetching current user details
		hiveIDs, ok := c.Locals("hive_ids").([]uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Checking if current user is authorized to perform this action
		authorized := slices.Contains(hiveIDs, targetHiveID)
		if !authorized {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Fetch chores belonging to the hive
		var chores []database.Chore
		err = db.Where("hive_id = ?", targetHiveID).Find(&chores).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"chores": chores})
	})

	router.Delete("/delete", func(c *fiber.Ctx) error {

		// Parsing and Validating
		idParam := c.Query("chore_id")

		parsed, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid chore_id")
		}

		targetChoreID := uint(parsed)

		// Fetch chore from DB
		var chore database.Chore
		err = db.First(&chore, targetChoreID).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "chore not found")
		}

		// Fetching current user details
		hiveIDs, ok := c.Locals("hive_ids").([]uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Checking if current user is authorized to perform this action
		authorized := slices.Contains(hiveIDs, chore.HiveID)
		if !authorized {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Delete chore
		err = db.Delete(&chore).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"msg": "deleted"})
	})
}
