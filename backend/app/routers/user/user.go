package user

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/storage"
	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRouter(router fiber.Router, db *gorm.DB, s3 *storage.Storage) {

	router.Get("/my-profile", func(c *fiber.Ctx) error {

		// Fetching UserID
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Get user details by ID
		user, err := utils.GetUserById(userID, db)
		if err != nil || user == nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Generating Display Image presigned URL
		url := ""

		if user.DisplayImgKey != nil && strings.TrimSpace(*user.DisplayImgKey) != "" {

			ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
			defer cancel()

			url, err = s3.DownloadPresignedURL(ctx, *user.DisplayImgKey, 10)

			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		// Building response
		res := myProfileResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			PhoneNumber:   user.PhoneNumber,
			DisplayImgURL: url,
			JoinedOn:      user.CreatedAt,
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"my_profile": res})
	})

	router.Post("/update-display-image", func(c *fiber.Ctx) error {

		// Fetching UserID
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		var reqBody updateUserDisplayImgSchema

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
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		imgKey := fmt.Sprintf("users/%d/display-image", userID)
		err = s3.Upload(ctx, imgKey, mimeType, data)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		//Update database
		err = db.Model(&database.User{}).Where("id = ?", userID).Update("display_img_key", imgKey).Error
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	router.Post("/search-by-phone", func(c *fiber.Ctx) error {

		var reqBody searchByPhoneSchema

		// Parsing and Validating
		err := utils.ParseAndValidate(c, &reqBody)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// Get user by phone number
		user, err := utils.GetUserByPhoneNumber(reqBody.PhoneNumber, db)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if user == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "user with this phone number not found"})
		}

		// Generating Display Image presigned URL
		url := ""

		if user.DisplayImgKey != nil && strings.TrimSpace(*user.DisplayImgKey) != "" {

			ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
			defer cancel()

			url, err = s3.DownloadPresignedURL(ctx, *user.DisplayImgKey, 10)
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		// Building response
		res := searchUserResponse{
			ID:            user.ID,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			PhoneNumber:   user.PhoneNumber,
			DisplayImgURL: url,
			JoinedOn:      user.CreatedAt,
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": res})
	})
}
