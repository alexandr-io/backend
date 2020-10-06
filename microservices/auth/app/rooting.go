package main

import (
	"github.com/alexandr-io/backend/auth/handlers"

	"github.com/gofiber/fiber"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {
	app.Post("/register", handlers.Register)
	//app.Post("/login", handlers.Login)
	//app.Post("/auth/refresh", handlers.RefreshAuthToken)
	//app.Get("/auth", userMiddleware.Protected(), handlers.Auth)

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})
}
