package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password *string) {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	*password = string(hashedBytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
