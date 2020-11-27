package main

import (
	"github.com/alexandr-io/backend/library/handlers"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createRoute(app *fiber.App) {
	// Recover middleware in case of panic
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

	app.Post("/library", userMiddleware.Protected(), handlers.LibraryCreation)
	app.Get("/library", userMiddleware.Protected(), handlers.LibraryRetrieve)
	app.Delete("/library", userMiddleware.Protected(), handlers.LibraryDelete)

	app.Get("/libraries", userMiddleware.Protected(), handlers.LibrariesRetrieve)

	app.Get("/book", userMiddleware.Protected(), handlers.BookRetrieve)
	app.Post("/book", userMiddleware.Protected(), handlers.BookCreation)
	app.Delete("/book", userMiddleware.Protected(), handlers.BookDelete)

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
