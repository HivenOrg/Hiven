package tests

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/HivenOrg/Hiven/config"
	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/start"
	"github.com/HivenOrg/Hiven/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type headers map[string]string
type payload map[string]any

var cfg config.Config
var testDB *gorm.DB
var testApp *fiber.App

func TestMain(m *testing.M) {

	ctx := context.Background()

	cfgPointer, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load test environment variables: %s", err.Error())
	}
	cfg = *cfgPointer

	if cfg.STAGE != "test" {
		log.Fatalf("STAGE must be 'test' for running tests, got: %s", cfg.STAGE)
	}

	testDB, err = database.ConnectToDB(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_SSLMODE, cfg.STAGE)
	if err != nil {
		log.Fatalf("failed to connect to test DB: %s", err.Error())
	}

	s3, err := storage.New(cfg.S3_BUCKET_NAME, cfg.AWS_REGION, cfg.AWS_ACCESS_KEY_ID, cfg.AWS_SECRET_ACCESS_KEY, cfg.STAGE, ctx)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	testApp = start.BuildApp(cfg, testDB, s3)

	code := m.Run()
	os.Exit(code)
}
