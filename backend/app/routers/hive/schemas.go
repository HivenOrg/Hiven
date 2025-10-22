package hive

import "mime/multipart"

type CreateHiveSchema struct {
	Name    string `json:"name" validate:"required,min=1"`
	Address string `json:"address" validate:"required,min=1"`
}

type UploadHiveDisplayImage struct {
	HiveID uint                  `form:"hive_id" validate:"required,gt=0"`
	File   *multipart.FileHeader `form:"file" validate:"required"`
}
