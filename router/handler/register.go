package handler

import (
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/model"
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

func Register(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if errors := user.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	uppercaseWithTimestamp(user)

	util.GenerateHashPassword(&user.Password)

	res, err := database.UserCollection.InsertOne(c.Context(), user)
	if err != nil {
		util.Logger("ERROR", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	id := res.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id.Hex()})
}

func uppercaseWithTimestamp(user *model.User) {
	user.Email = strings.ToUpper(user.Email)
	user.FirstName = strings.ToUpper(user.FirstName)
	user.LastName = strings.ToUpper(user.LastName)
	user.Role = strings.ToUpper(user.Role)
	user.CreatedAt = time.Now()
}
