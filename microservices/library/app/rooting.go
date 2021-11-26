package main

import (
	"github.com/alexandr-io/backend/library/handlers"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

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

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	handlers.CreateUserLibraryHandlers(app)
	handlers.CreateLibraryHandlers(app)
	handlers.CreateBookHandlers(app)
	handlers.CreateBookProgressHandlers(app)
	handlers.CreateUserDataHandlers(app)
	handlers.CreateGroupHandlers(app)
	handlers.CreatePermissionHandlers(app)
	handlers.CreateMetadataHandlers(app)
	handlers.CreateProgressSpeedHandlers(app)

	// Retrieve definitions from dictionary API
	app.Get("/dictionary/definition/:lang/:queried_word", userMiddleware.Protected(), handlers.DictionaryRetrieve)

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
