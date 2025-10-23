package utils

import (
	"errors"
	"slices"

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

// Checking if current user is active owner
func IsActiveOwner(currUserID uint, currUserHiveIDs []uint, targetHiveID uint, db *gorm.DB) (bool, error) {

	authorized := slices.Contains(currUserHiveIDs, targetHiveID)
	if !authorized {
		return false, nil
	}

	member := database.Member{}
	err := db.Where("hive_id = ? and user_id = ?", targetHiveID, currUserID).First(&member).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	authorized = member.Status == "active" && member.Role == "owner"
	if !authorized {
		return false, nil
	}

	return true, nil
}
