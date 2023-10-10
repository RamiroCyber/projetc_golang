package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvironment() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	port := os.Getenv("port_application")

	return port
}
