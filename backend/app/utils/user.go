package utils

import (
	"errors"

	"github.com/HivenOrg/Hiven/database"
	"gorm.io/gorm"
)

func getUserByField(field string, value any, db *gorm.DB) (*database.User, error) {

	var user database.User
	result := db.Where(field+" = ?", value).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // user not found
		}
		return nil, result.Error // some other db error
	}
	return &user, nil // user found
}

func GetUserById(id uint, db *gorm.DB) (*database.User, error) {
	return getUserByField("id", id, db)
}

func GetUserByEmail(email string, db *gorm.DB) (*database.User, error) {
	return getUserByField("email", email, db)
}

func GetUserByPhoneNumber(phoneNumber string, db *gorm.DB) (*database.User, error) {
	return getUserByField("phone_number", phoneNumber, db)
}

func GetHiveIDsForUser(userID uint, db *gorm.DB) ([]uint, error) {

	var hiveIDs []uint

	err := db.Model(&database.Member{}).Where("user_id = ?", userID).Pluck("hive_id", &hiveIDs).Error

	if err != nil {
		return nil, err
	}

	return hiveIDs, nil
}
