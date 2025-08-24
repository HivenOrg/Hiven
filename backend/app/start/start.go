package start

import "github.com/gofiber/fiber/v2"

/*
Loading environment variables and connecting to database is seperated from setting up the fiber app
This is done to make automated testing possible
*/

func Server() *fiber.App {
	app := BuildApp()
	return app
}

func BuildApp() *fiber.App {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Hiven API!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("The server is running :)")
	})

	return app
}
