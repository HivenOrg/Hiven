package tests

import (
	"log"
	"os"
	"testing"

	"github.com/HivenOrg/Hiven/config"
	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/start"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type headers map[string]string
type payload map[string]any

var cfg config.Config
var testDB *gorm.DB
var testApp *fiber.App

func TestMain(m *testing.M) {

	var err error

	cfgPointer, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load test environment variables: %s", err.Error())
	}
	cfg = *cfgPointer

	testDB, err = database.ConnectToDB(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_SSLMODE, cfg.STAGE)
	if err != nil {
		log.Fatalf("failed to connect to test DB: %s", err.Error())
	}

	testApp = start.BuildApp(cfg, testDB)

	code := m.Run()
	os.Exit(code)
}
