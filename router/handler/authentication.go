package handler

import (
	"errors"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/model"
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/RamiroCyber/projetc_golang/util/constants"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func Login(c *fiber.Ctx) error {
	auth := new(model.Authentication)

	if err := c.BodyParser(auth); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Cannot parse JSON")
	}

	if errs := auth.Validate(); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	user := new(model.User)
	if err := database.UserCollection.FindOne(c.Context(), bson.M{"email": strings.ToUpper(auth.Email)}).Decode(&user); errors.Is(err, mongo.ErrNoDocuments) {
		return c.Status(fiber.StatusUnauthorized).SendString(constants.InvalidAuth)
	} else if err != nil {
		util.Logger(constants.Error, err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Database error")
	}

	if err := util.CheckPasswordHash(auth.Password, user.Password); err == false {
		return c.Status(fiber.StatusUnauthorized).SendString(constants.InvalidAuth)
	}

	tokenString, err := config.GenerateToken(user.ID.Hex(), c)
	if err != nil {
		util.Logger(constants.Error, err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
}
