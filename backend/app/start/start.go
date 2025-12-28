package start

import (
	"context"
	"log"

	"github.com/HivenOrg/Hiven/config"
	"github.com/HivenOrg/Hiven/database"
	"github.com/HivenOrg/Hiven/middleware"
	"github.com/HivenOrg/Hiven/routers/auth"
	"github.com/HivenOrg/Hiven/routers/chore"
	"github.com/HivenOrg/Hiven/routers/hive"
	"github.com/HivenOrg/Hiven/routers/shopping"
	"github.com/HivenOrg/Hiven/routers/user"
	"github.com/HivenOrg/Hiven/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
Loading environment variables and connecting to database is seperated from setting up the fiber app
This is done to make automated testing possible
*/
func Server() *fiber.App {

	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	db, err := database.ConnectToDB(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_SSLMODE, cfg.STAGE)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	s3, err := storage.New(cfg.S3_BUCKET_NAME, cfg.AWS_REGION, cfg.AWS_ACCESS_KEY_ID, cfg.AWS_SECRET_ACCESS_KEY, cfg.STAGE, ctx)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	return BuildApp(*cfg, db, s3)
}

func BuildApp(cfg config.Config, db *gorm.DB, s3 *storage.Storage) *fiber.App {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Hiven API!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("The server is running :)")
	})

	// Adding auth routes
	auth.AuthRouter(app.Group("/auth"), db, cfg)

	// Protecting non-auth routes
	app.Use(middleware.AuthMiddleware(db, cfg))

	app.Get("/test/protected", func(c *fiber.Ctx) error {
		return c.SendString("Token is valid")
	})

	hive.HiveRouter(app.Group("/hive"), db, s3)
	user.UserRouter(app.Group("/user"), db, s3)
	chore.ChoreRouter(app.Group("/chore"), db)
	shopping.ShoppingRouter(app.Group("/shopping"), db)

	return app
}
