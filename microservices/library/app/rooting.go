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

	app.Get("/metadata", userMiddleware.Protected(), handlers.MetadataRetrieve)

	app.Get("/libraries", userMiddleware.Protected(), handlers.LibrariesRetrieve)

	// Library relatives URLs
	app.Post("/library", userMiddleware.Protected(), handlers.LibraryCreation)
	app.Get("/library/:library_id", userMiddleware.Protected(), handlers.LibraryRetrieve)
	app.Delete("/library/:library_id", userMiddleware.Protected(), handlers.LibraryDelete)

	// Books relatives URLs
	app.Get("/library/:library_id/books", userMiddleware.Protected(), handlers.BooksRetrieve)
	app.Post("/library/:library_id/book", userMiddleware.Protected(), handlers.BookCreation)
	app.Get("/library/:library_id/book/:book_id", userMiddleware.Protected(), handlers.BookRetrieve)
	app.Post("/library/:library_id/book/:book_id", userMiddleware.Protected(), handlers.BookUpdate)
	app.Delete("/library/:library_id/book/:book_id", userMiddleware.Protected(), handlers.BookDelete)

	// Book progress relatives URLs
	app.Get("/library/:library_id/book/:book_id/progress", userMiddleware.Protected(), handlers.ProgressRetrieve)
	app.Post("/library/:library_id/book/:book_id/progress", userMiddleware.Protected(), handlers.ProgressUpsert)

	// Userdata relative URLs
	app.Get("/library/:library_id/book/:book_id/data", userMiddleware.Protected(), handlers.UserDataList)
	app.Post("/library/:library_id/book/:book_id/data", userMiddleware.Protected(), handlers.UserDataCreate)
	app.Delete("/library/:library_id/book/:book_id/data", userMiddleware.Protected(), handlers.UserDataDeleteAllIn)
	app.Get("/library/:library_id/book/:book_id/data/:data_id", userMiddleware.Protected(), handlers.UserDataGet)
	app.Post("/library/:library_id/book/:book_id/data/:data_id", userMiddleware.Protected(), handlers.UserDataUpdate)
	app.Delete("/library/:library_id/book/:book_id/data/:data_id", userMiddleware.Protected(), handlers.UserDataDeleteOne)

	// Permissions relative URLs
	//
	// Group relative URLs
	app.Post("/library/:library_id/group", userMiddleware.Protected(), handlers.GroupCreate)
	app.Get("/library/:library_id/group/:group_id", userMiddleware.Protected(), handlers.GroupRetrieve)
	app.Post("/library/:library_id/group/:group_id", userMiddleware.Protected(), handlers.GroupUpdate)
	app.Delete("/library/:library_id/group/:group_id", userMiddleware.Protected(), handlers.GroupDelete)

	app.Post("/library/:library_id/group/:group_id/join", userMiddleware.Protected(), handlers.GroupAddUser)
	app.Get("/library/:library_id/user/groups", userMiddleware.Protected(), handlers.GroupsRetrieveUser)

	app.Get("/library/:library_id/permissions", userMiddleware.Protected(), handlers.UserLibraryPermissionRetrieve)

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
