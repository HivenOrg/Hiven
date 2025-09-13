package database

import "time"

type User struct {
	ID            uint `gorm:"primaryKey"`
	Email         string
	Password      string
	FirstName     string
	LastName      string
	PhoneNumber   string // store in E.164 format
	DisplayImgKey *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Hive struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Address       string
	DisplayImgKey *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Member struct {
	HiveID   uint `gorm:"primaryKey"`
	UserID   uint `gorm:"primaryKey"`
	Status   string
	Role     string
	JoinedAt time.Time
}
