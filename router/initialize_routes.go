package router

import (
	"fmt"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/RamiroCyber/projetc_golang/middleware"
	"github.com/RamiroCyber/projetc_golang/router/handler"
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes() *fiber.App {
	app := config.ConfigsRoutes()

	api := app.Group(util.API)

	v1 := api.Group(fmt.Sprint("/", util.V1), func(c *fiber.Ctx) error {
		c.Set(util.VERSION, util.V1)
		return c.Next()
	})

	//HEALTHCHECK
	v1.Get("/health", handler.HealthCheck)

	//AUTHENTICATION
	v1.Post("/login", handler.Login)

	//USER
	v1.Post("/user", middleware.JWTMiddleware, handler.CreateUser)

	return app
}
