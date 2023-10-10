package routes

import (
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/gofiber/fiber/v2"
)

func Routes() *fiber.App {
	app := config.ConfigsRoutes()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return app
}
