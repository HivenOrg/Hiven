package start

import (
	"log"

	"github.com/HivenOrg/Hiven/config"
	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/routers/auth"
	"github.com/gofiber/fiber/v2"
)

func BuildApp() *fiber.App {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	db, err := database.ConnectToDB(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_SSLMODE, cfg.STAGE)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

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
