package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}
