package database

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	Firstname   string    `gorm:"not null"`
	Lastname    string    `gorm:"not null"`
	PhoneNumber string    `gorm:"uniqueIndex;not null"` // store in E.164 format
	CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:current_timestamp"`
}
