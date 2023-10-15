package config

import (
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(userId string, c *fiber.Ctx) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		Issuer:    "myApp",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		util.Logger("ERROR", err.Error())
		return "", err
	}
	c.Set("Authorization", "Bearer "+tokenString)

	return tokenString, nil
}
