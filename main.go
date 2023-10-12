package main

import (
	"fmt"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/router"
)

func main() {
	port := config.LoadEnvironment()

	database.Connect()
	database.SetupCollections()

	app := router.InitializeRoutes()

	app.Listen(fmt.Sprintf(":%s", port))

}
