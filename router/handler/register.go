package handler

import (
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/model"
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/RamiroCyber/projetc_golang/util/constants"
	"github.com/gofiber/fiber/v2"
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

	prepareUserData(user)

	if !util.IsValidPhoneNumber(user.Phone) {
		return c.Status(fiber.StatusBadRequest).SendString(constants.InvalidPhone)
	}

	if err := saveUserToDatabase(c, user); err != nil {
		util.Logger(constants.Error, err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(constants.FailedRegister)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func prepareUserData(user *model.User) {
	uppercaseFields(user)
	setTimestamp(user)
	util.GenerateHashPassword(&user.Password)
}

func uppercaseFields(user *model.User) {
	user.Email = strings.ToUpper(user.Email)
	user.FirstName = strings.ToUpper(user.FirstName)
	user.LastName = strings.ToUpper(user.LastName)
	user.Role = strings.ToUpper(user.Role)
}

func setTimestamp(user *model.User) {
	user.CreatedAt = time.Now()
}

func saveUserToDatabase(c *fiber.Ctx, user *model.User) error {
	_, err := database.UserCollection.InsertOne(c.Context(), user)
	return err
}
