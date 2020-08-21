package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	createRoute(app)

	app.Listen(3000)
}
