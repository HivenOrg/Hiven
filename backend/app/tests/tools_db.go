package tests

import (
	"fmt"

	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/utils"
	"gorm.io/gorm"
)

func createDummyUser(testDB *gorm.DB) (*database.User, error) {

	plainPassword := "iwillfindit"
	hashedPassword, err := utils.HashPassword(plainPassword)

	if err != nil {
		return nil, err
	}

	user := &database.User{
		Email:       "timtracer@test.com",
		Password:    hashedPassword,
		FirstName:   "Tim",
		LastName:    "Tracer",
		PhoneNumber: "+10000000000",
	}

	if err := testDB.Create(user).Error; err != nil {
		return nil, err
	}

	user.Password = plainPassword //returning user with plain password

	return user, nil
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
