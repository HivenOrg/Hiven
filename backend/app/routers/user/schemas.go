package user

import (
	"mime/multipart"
	"time"
)

type myProfileResponse struct {
	ID            uint      `json:"id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	DisplayImgURL string    `json:"display_img_url"`
	JoinedOn      time.Time `json:"joined_on"`
}

type updateUserDisplayImgSchema struct {
	File *multipart.FileHeader `form:"file" validate:"required"`
}

type searchByPhoneSchema struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"` // international telephone numbering standard
}

type searchUserResponse struct {
	ID            uint      `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	DisplayImgURL string    `json:"display_img_url"`
	JoinedOn      time.Time `json:"joined_on"`
}
