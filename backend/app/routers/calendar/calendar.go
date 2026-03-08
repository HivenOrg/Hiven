package calendar

import (
	"slices"
	"strconv"
	"time"

	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CalendarRouter(router fiber.Router, db *gorm.DB) {

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

		// Fetch events belonging to the hive
		var events []database.CalendarEvent
		err = db.Where("hive_id = ?", targetHiveID).Order("created_at DESC").Find(&events).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"events": events})
	})

	router.Post("/create", func(c *fiber.Ctx) error {

		var reqBody createEvent

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

		// Parsing string into time.time
		eventTime, err := time.Parse(time.RFC3339, reqBody.EventTimestamp)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid event_timestamp format")
		}

		// Adding new calendar event to database
		newEvent := database.CalendarEvent{
			HiveID:              reqBody.HiveID,
			CreatorID:           userID,
			EventTitle:          reqBody.EventTitle,
			EventTimestamp:      eventTime,
			EventOriginTimezone: reqBody.EventOriginTimezone,
		}

		err = db.Create(&newEvent).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(newEvent)
	})

	router.Delete("/delete", func(c *fiber.Ctx) error {

		// Parsing and Validating
		idParam := c.Query("event_id")

		parsed, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid event_id")
		}

		targetEventID := uint(parsed)

		// Fetching current user details
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		hiveIDs, ok := c.Locals("hive_ids").([]uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Fetch Event from DB
		var event database.CalendarEvent
		err = db.First(&event, targetEventID).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "item not found")
		}

		// Checking if current user is authorized to perform this action
		authorized := slices.Contains(hiveIDs, event.HiveID)
		if !authorized || event.CreatorID != userID {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Delete item
		err = db.Delete(&event).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"msg": "deleted"})
	})
}
