package tests

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/storage"
	"github.com/HivenOrg/Hiven/utils"
	"gorm.io/gorm"
)

func createDummyUser(testDB *gorm.DB, jwtSecretKey string) (*database.User, string, error) {

	plainPassword := "iwillfindit"
	hashedPassword, err := utils.HashPassword(plainPassword)

	if err != nil {
		return nil, "", err
	}

	suffix := fmt.Sprintf("%d", time.Now().UnixNano()+rand.Int63n(1000))
	email := fmt.Sprintf("timtracer_%s@test.com", suffix)
	phone_number := fmt.Sprintf("+1%09d", rand.Intn(1000000000))

	user := &database.User{
		Email:       email,
		Password:    hashedPassword,
		FirstName:   "Tim",
		LastName:    "Tracer",
		PhoneNumber: phone_number,
	}

	if err := testDB.Create(user).Error; err != nil {
		return nil, "", err
	}

	user.Password = plainPassword //returning user with plain password

	token, err := utils.GenerateJWT(user.ID, jwtSecretKey, 1)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func createDummyHive(testDB *gorm.DB, userID uint) (*database.Hive, error) {

	hive := &database.Hive{
		Name:    "London Home",
		Address: "48 Camden High Street, Camden Town, London",
	}

	if err := testDB.Create(hive).Error; err != nil {
		return nil, err
	}

	member := &database.Member{
		HiveID: hive.ID,
		UserID: userID,
		Status: "active",
		Role:   "owner",
	}

	if err := testDB.Create(member).Error; err != nil {
		return nil, err
	}

	return hive, nil
}

func resetTestDB(testDB *gorm.DB, stage string) error {

	models := []any{
		&database.User{},
		&database.Hive{},
		&database.Member{},
	}

	err := testDB.Migrator().DropTable(models...)
	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	err = database.CreateTables(testDB, stage)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}

func resetTestStorage(testStorage *storage.Storage) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// List all object keys in the test bucket
	keys, err := testStorage.ListObjects(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to list test bucket objects: %w", err)
	}

	// Delete each object
	for _, key := range keys {
		if err := testStorage.DeleteObject(ctx, key); err != nil {
			return fmt.Errorf("failed to delete object %q: %w", key, err)
		}
	}

	return nil
}
