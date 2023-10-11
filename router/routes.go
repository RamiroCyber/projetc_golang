package router

import (
	"fmt"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/gofiber/fiber/v2"
)

const (
	V1  = "v1"
	API = "/api"
)

func InitializeRoutes() *fiber.App {
	app := config.ConfigsRoutes()

	api := app.Group(API)

	v1 := api.Group(fmt.Sprint("/", V1), func(c *fiber.Ctx) error {
		c.Set("Version", V1)
		return c.Next()
	})

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return app
}
