package database

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	PhoneNumber string    `gorm:"uniqueIndex;not null"` // store in E.164 format
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
