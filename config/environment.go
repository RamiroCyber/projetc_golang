package config

import (
	"fmt"
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		util.Logger("ERROR", fmt.Sprintf(".env: %v", err))
	}
}
