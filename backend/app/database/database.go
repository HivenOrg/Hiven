package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(host, user, password, dbname, port, sslmode, stage string) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = ping(db)
	if err != nil {
		return nil, err
	}

	if stage == "dev" {
		err = db.AutoMigrate(&User{})
		if err != nil {
			return nil, fmt.Errorf("failed to auto-migrate tables: %w", err)
		}
	}

	return db, nil
}

func ping(db *gorm.DB) error {

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}
