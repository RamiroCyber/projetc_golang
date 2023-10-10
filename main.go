package main

import "github.com/RamiroCyber/projetc_golang/routes"

func main() {
	app := routes.Routes()

	app.Listen(":5000")
}
