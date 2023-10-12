package handler

import (
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/models"
	"github.com/RamiroCyber/projetc_golang/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if errors := user.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	utils.GenerateHashPassword(&user.Password)

	user.CreatedAt = time.Now()

	res, err := database.UserCollection.InsertOne(c.Context(), user)
	if err != nil {
		utils.Logger("ERROR", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	id := res.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id.Hex()})
}
