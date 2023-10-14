package handler

import (
	"errors"
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/model"
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	authentication := new(model.Authentication)

	if err := c.BodyParser(authentication); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Cannot parse JSON")
	}

	if errs := authentication.Validate(); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	user := new(model.User)
	if err := database.UserCollection.FindOne(c.Context(), bson.M{"email": authentication.Email}).Decode(&user); errors.Is(err, mongo.ErrNoDocuments) {
		return c.Status(fiber.StatusUnauthorized).SendString("User not found")
	} else if err != nil {
		util.Logger("ERROR", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Database error")
	}

	if err := util.CheckPasswordHash(authentication.Password, user.Password); err == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid password"})
	}

	claims := &jwt.RegisteredClaims{
		Subject:   user.ID.Hex(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error generating token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
}
