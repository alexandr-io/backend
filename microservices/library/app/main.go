package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Library Service started")

	// Fiber
	app := fiber.New()
	createRoute(app)
	log.Fatal(app.Listen(":3000"))
}
