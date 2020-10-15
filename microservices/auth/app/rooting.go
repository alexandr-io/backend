package main

import (
	"github.com/alexandr-io/backend/auth/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {
	// Recover middleware in case of panic
	app.Use(recover.New())
	//15 Oct 09:36:56 CEST | 504 |   523ms |      172.19.0.1 | POST    | /register        | {"message":"Kafka register response timed out","file":"/app/kafka/register.go","line":143,"content-type":"text/html; charset=utf-8","custom-message":""}
	app.Use(logger.New(logger.Config{
		TimeFormat: "2 Jan 15:04:05 MST",
		TimeZone:   "Europe/Paris",
		Next: func(ctx *fiber.Ctx) bool {
			if string(ctx.Request().RequestURI()) == "/register" {
				return false
			}
			return true
		},
	}))
	app.Get("/dashboard", monitor.New())

	app.Post("/register", handlers.Register)
	//app.Post("/login", handlers.Login)
	//app.Post("/auth/refresh", handlers.RefreshAuthToken)
	//app.Get("/auth", userMiddleware.Protected(), handlers.Auth)

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}