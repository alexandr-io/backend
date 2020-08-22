package main

import (
	"github.com/Alexandr-io/Backend/User/handlers"
	"github.com/gofiber/fiber"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {

	app.Post("/register", handlers.Register)

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})
}
