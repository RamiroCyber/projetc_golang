package router

import (
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes() *fiber.App {
	app := config.ConfigsRoutes()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return app
}
