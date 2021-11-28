package main

import (
	"time"

	"github.com/alexandr-io/backend/user/handlers"
	userMiddleware "github.com/alexandr-io/backend/user/middleware"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// createRoute creates all the routes of the service.
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
	app.Use(limiter.New(limiter.Config{
		Max:				30,	
		Expiration:			15 * time.Second,
		LimitReached:		func(c *fiber.Ctx) error{
			return c.SendString("Rate limited, your IP is sending too many requests")
		},
	}))
	app.Get("/dashboard", monitor.New())

	app.Get("/user", userMiddleware.Protected(), handlers.GetUser)
	app.Put("/user", userMiddleware.Protected(), handlers.UpdateUser)
	app.Delete("/user", userMiddleware.Protected(), handlers.DeleteUser)

	app.Get("/verify", handlers.VerifyEmail)
	app.Get("/verify/update", handlers.VerifyUpdateEmail)
	app.Get("/email/cancel", handlers.CancelEmailUpdate)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
