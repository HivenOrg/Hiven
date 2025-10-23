package hive

import (
	"mime/multipart"
	"time"
)

type createHiveSchema struct {
	Name    string `json:"name" validate:"required,min=1"`
	Address string `json:"address" validate:"required,min=1"`
}

type uploadHiveDisplayImage struct {
	HiveID uint                  `form:"hive_id" validate:"required,gt=0"`
	File   *multipart.FileHeader `form:"file" validate:"required"`
}

type hiveResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	DisplayImgURL string    `json:"display_img_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
