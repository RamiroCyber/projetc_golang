package router

import (
	"fmt"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/RamiroCyber/projetc_golang/middleware"
	"github.com/RamiroCyber/projetc_golang/router/handler"
	"github.com/RamiroCyber/projetc_golang/util/constants"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes() *fiber.App {
	app := config.ConfigsRoutes()

	api := app.Group(constants.API)

	v1 := api.Group(fmt.Sprint("/", constants.V1), func(c *fiber.Ctx) error {
		c.Set(constants.VERSION, constants.V1)
		return c.Next()
	})

	//HEALTHCHECK
	v1.Get("/health", handler.HealthCheck)

	//AUTHENTICATION
	v1.Post("/login", handler.Login)

	//USER
	v1.Post("/register", handler.Register)
	v1.Delete("/user/:id", middleware.JWTMiddleware, handler.DeleteUser)

	return app
}
