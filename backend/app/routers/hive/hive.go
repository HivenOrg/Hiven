package hive

import (
	"context"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/storage"
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HiveRouter(router fiber.Router, db *gorm.DB, s3 *storage.Storage) {

	router.Post("/create", func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
		defer cancel()

		var reqBody createHiveSchema

		// Parsing and Validating
		err := utils.ParseAndValidate(c, &reqBody)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// Fetching UserID
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Creating a new Hive, adding current user as owner, generating presigned url for display picture
		newHive := &database.Hive{
			Name:    reqBody.Name,
			Address: reqBody.Address,
		}

		err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			// Adding new hive to DB
			if err := tx.Create(newHive).Error; err != nil {
				return err
			}

			// Adding current user as owner
			newMember := &database.Member{
				HiveID: newHive.ID,
				UserID: userID,
				Status: "active",
				Role:   "owner",
			}
			if err := tx.Create(newMember).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"hive_id": newHive.ID})
	})

	router.Post("/update-display-image", func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
		defer cancel()

		var reqBody uploadHiveDisplayImage

		// Parse hive_id from form
		hiveIDStr := c.FormValue("hive_id")
		hiveID, err := strconv.ParseUint(hiveIDStr, 10, 64)
		if err != nil || hiveID == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "invalid or missing hive_id")
		}
		reqBody.HiveID = uint(hiveID)

		// Parse file from form
		file, err := c.FormFile("file")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "missing file")
		}
		reqBody.File = file

		// Checking upload file type
		mimeType := storage.InferMimeTypeFromFilename(reqBody.File.Filename)
		if !(mimeType == "image/jpeg" || mimeType == "image/png") {
			return fiber.NewError(fiber.StatusBadRequest, "only JPG/JPEG and PNG images are allowed")
		}

		// Fetching HiveIDs
		hiveIDs, ok := c.Locals("hive_ids").([]uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Checking if current user is authorized to perform this action on this Hive
		authorized := slices.Contains(hiveIDs, reqBody.HiveID)
		if !authorized {
			return c.SendStatus(fiber.StatusForbidden)
		}

		// Converting file to bytes
		src, err := file.Open()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer src.Close()

		data, err := io.ReadAll(src)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Uploading to S3
		imgKey := fmt.Sprintf("hives/%d/display-image", reqBody.HiveID)
		err = s3.Upload(ctx, imgKey, mimeType, data)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		//Update database
		err = db.Model(&database.Hive{}).Where("id = ?", reqBody.HiveID).Update("display_img_key", imgKey).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	router.Get("/my-hives", func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
		defer cancel()

		// Fetching HiveIDs
		hiveIDs, ok := c.Locals("hive_ids").([]uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if len(hiveIDs) == 0 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"my_hives": []hiveResponse{}})
		}

		// Fetching hives from database
		var hives []database.Hive
		err := db.Where("id IN ?", hiveIDs).Find(&hives).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Parsing response
		var res []hiveResponse

		for _, h := range hives {

			// Generating presigned urlS
			url := ""
			if h.DisplayImgKey != nil && strings.TrimSpace(*h.DisplayImgKey) != "" {
				url, err = s3.DownloadPresignedURL(ctx, *h.DisplayImgKey, 10)
				if err != nil {
					return c.SendStatus(fiber.StatusInternalServerError)
				}
			}

			res = append(res, hiveResponse{
				ID:            h.ID,
				Name:          h.Name,
				Address:       h.Address,
				DisplayImgURL: url,
				CreatedAt:     h.CreatedAt,
				UpdatedAt:     h.UpdatedAt,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"my_hives": res})
	})
}
