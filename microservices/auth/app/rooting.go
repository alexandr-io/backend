package main

import (
	"github.com/alexandr-io/backend/auth/handlers"
	authMiddleware "github.com/alexandr-io/backend/auth/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {
	// Recover middleware in case of panic
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2 Jan 15:04:05 MST",
		TimeZone:   "Europe/Paris",
		Next: func(ctx *fiber.Ctx) bool {
			if string(ctx.Request().RequestURI()) == "/dashboard" ||
				string(ctx.Request().RequestURI()) == "/docs" ||
				string(ctx.Request().RequestURI()) == "/swagger.yml" {
				return true
			}
			return false
		},
	}))
	app.Get("/dashboard", monitor.New())

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/refresh", handlers.RefreshAuthToken)
	app.Get("/auth", authMiddleware.Protected(), handlers.Auth)
	app.Post("/logout", authMiddleware.Protected(), handlers.Logout)

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
