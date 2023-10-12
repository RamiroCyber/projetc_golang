package main

import (
	"context"
	"fmt"
	"github.com/RamiroCyber/projetc_golang/config"
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/router"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadEnvironment()

	database.ConnectDB()
	database.SetupCollections()

	app := router.InitializeRoutes()

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("port_application"))); err != nil {
			log.Panicf("Falha ao iniciar o servidor : %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	database.Client.Disconnect(ctx)

	log.Println("Servidor Pausado")

}
