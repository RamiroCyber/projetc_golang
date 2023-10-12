package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password *string) {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	*password = string(hashedBytes)
}
