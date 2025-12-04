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
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

type Hive struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Address       string
	DisplayImgKey *string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

type Member struct {
	HiveID   uint `gorm:"primaryKey"`
	UserID   uint `gorm:"primaryKey"`
	Status   string
	Role     string
	JoinedAt time.Time `gorm:"autoCreateTime"`
}

type Chore struct {
	ID        uint `gorm:"primaryKey"`
	HiveID    uint
	CreatorID uint
	Chore     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
