package start

import (
	"log"

	"github.com/HivenOrg/Hiven/config"
	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/routers/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
Loading environment variables and connecting to database is seperated from setting up the fiber app
This is done to make automated testing possible
*/
func Server() *fiber.App {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	db, err := database.ConnectToDB(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_SSLMODE, cfg.STAGE)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	return BuildApp(cfg, db)
}

func BuildApp(cfg *config.Config, db *gorm.DB) *fiber.App {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Hiven API!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("The server is running :)")
	})

	// Adding auth routes
	auth.AuthRouter(app.Group("/auth"), db, cfg)

	return app
}
