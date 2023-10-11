package router

import (
	"fmt"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/RamiroCyber/projetc_golang/router/handles"
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
	v1.Get("/health", handles.HealthCheck)

	//OPENING
	v1.Get("/opportunities", handles.HealthCheck)
	v1.Post("/opportunities", handles.HealthCheck)
	v1.Put("/opportunities", handles.HealthCheck)
	v1.Delete("/opportunities", handles.HealthCheck)

	return app
}
