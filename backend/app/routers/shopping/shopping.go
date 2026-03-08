package shopping

import (
	"slices"
	"strconv"

	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShoppingRouter(router fiber.Router, db *gorm.DB) {

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
		var items []database.ShoppingItem
		err = db.Where("hive_id = ?", targetHiveID).Order("created_at DESC").Find(&items).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"items": items})
	})

	router.Post("/add", func(c *fiber.Ctx) error {

		var reqBody addItem

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

		// Adding new shopping item to database
		newItem := database.ShoppingItem{
			HiveID:    reqBody.HiveID,
			CreatorID: userID,
			Item:      reqBody.Item,
		}

		err = db.Create(&newItem).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(newItem)
	})

	router.Delete("/remove", func(c *fiber.Ctx) error {

		// Parsing and Validating
		idParam := c.Query("item_id")

		parsed, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid item_id")
		}

		targetItemID := uint(parsed)

		// Fetch item from DB
		var item database.ShoppingItem
		err = db.First(&item, targetItemID).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "item not found")
		}

		// Fetching current user details
		hiveIDs, ok := c.Locals("hive_ids").([]uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Checking if current user is authorized to perform this action
		authorized := slices.Contains(hiveIDs, item.HiveID)
		if !authorized {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Delete item
		err = db.Delete(&item).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"msg": "deleted"})
	})
}
