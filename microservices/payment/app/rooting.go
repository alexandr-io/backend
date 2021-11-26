package main

import (
	"github.com/alexandr-io/backend/payment/handlers"
	"github.com/alexandr-io/backend/payment/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createRoute(app *fiber.App) {
	// Recover middleware in case of panic
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2 Jan 15:04:05 MST",
		TimeZone:   "Europe/Paris",
		Next: func(ctx *fiber.Ctx) bool {
			if string(ctx.Request().RequestURI()) == "/dashboard" {
				return true
			}
			return false
		},
	}))

	app.Get("/subscriptions", handlers.ListSubscriptions)
	app.Post("/subscribe", middleware.Protected(), handlers.Subscribe)
	app.Get("/subscription/price", middleware.Protected(), handlers.GetCustomerSubscription)

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
