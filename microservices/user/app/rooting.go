package main

import "github.com/gofiber/fiber"

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {

	app.Post("/register", register)

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})
}
