package utils

import (
	"errors"

	"github.com/HivenOrg/Hiven/database"
	"gorm.io/gorm"
)

func UserWithEmailExists(email string, db *gorm.DB) (bool, error) {

	var user database.User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // user not found
		}
		return false, result.Error // some other db error
	}

	return true, nil // user found
}

func GetUserByEmail(email string, db *gorm.DB) (*database.User, error) {

	var user database.User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // user not found
		}
		return nil, result.Error // some other db error
	}

	return &user, nil // user found
}

func UserWithPhoneNumberExists(phoneNumber string, db *gorm.DB) (bool, error) {

	var user database.User
	result := db.Where("phone_number = ?", phoneNumber).First(&user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // user not found
		}
		return false, result.Error // some other db error
	}

	return true, nil // user found
}

func GetHiveIDsForUser(userID uint, db *gorm.DB) ([]uint, error) {

	var hiveIDs []uint

	err := db.Model(&database.Member{}).Where("user_id = ?", userID).Pluck("hive_id", &hiveIDs).Error

	if err != nil {
		return nil, err
	}

	return hiveIDs, nil
}
