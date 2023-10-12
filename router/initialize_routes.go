package router

import (
	"fmt"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/RamiroCyber/projetc_golang/router/handler"
	"github.com/RamiroCyber/projetc_golang/utils"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes() *fiber.App {
	app := config.ConfigsRoutes()

	api := app.Group(utils.API)

	v1 := api.Group(fmt.Sprint("/", utils.V1), func(c *fiber.Ctx) error {
		c.Set(utils.VERSION, utils.V1)
		return c.Next()
	})

	//HEALTHCHECK
	v1.Get("/health", handler.HealthCheck)

	//USER
	v1.Post("/user", handler.CreateUser)

	return app
}
