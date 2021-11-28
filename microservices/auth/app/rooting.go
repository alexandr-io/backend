package main

import (
	"time"

	"github.com/alexandr-io/backend/auth/handlers"
	authMiddleware "github.com/alexandr-io/backend/auth/middleware"
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

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/refresh", handlers.RefreshAuthToken)
	app.Get("/auth", authMiddleware.Protected(), handlers.Auth)
	app.Post("/logout", authMiddleware.Protected(), handlers.Logout)
	app.Post("/password/reset", handlers.SendResetPasswordEmail)
	app.Get("/password/reset", handlers.ResetPasswordInfoFromToken)
	app.Put("/password/reset", handlers.ResetPassword)
	app.Put("/password/update", authMiddleware.Protected(), handlers.UpdatePassword)
	app.Get("/invitation/new", handlers.NewInvitation)
	app.Delete("/invitation/:token", authMiddleware.Protected(), handlers.DeleteInvitation)

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
