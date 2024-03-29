package main

import (
	"net/http"

	"github.com/alexandr-io/backend/media/handlers"
	mediaMiddleware "github.com/alexandr-io/backend/media/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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
	app.Get("/dashboard", monitor.New())

	app.Use(filesystem.New(filesystem.Config{Root: http.Dir("/media")}))

	app.Post("/book/upload", mediaMiddleware.Protected(), handlers.UploadBook)
	app.Get("/book/:book_id/download", mediaMiddleware.Protected(), handlers.DownloadBook)
	app.Post("/book/cover/upload", mediaMiddleware.Protected(), handlers.UploadBookCover)
	app.Get("/book/:book_id/cover", mediaMiddleware.Protected(), handlers.BookCover)

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
