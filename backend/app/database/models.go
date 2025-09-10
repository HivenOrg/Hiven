package database

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	PhoneNumber string    `gorm:"unique;not null"` // store in E.164 format
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}
