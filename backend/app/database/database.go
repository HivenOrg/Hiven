package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(host, user, password, dbname, port, sslmode string) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	err = ping(db)
	if err != nil {
		return nil, err
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
