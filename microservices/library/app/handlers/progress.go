package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	bookProgressServ "github.com/alexandr-io/backend/library/internal/bookprogress"
	permissionServ "github.com/alexandr-io/backend/library/internal/permission"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// progressUpsert updates the user's progression on a book
func progressUpsert(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Parse data
	var bookProgress data.BookProgressData
	if err := ParseBodyJSON(ctx, &bookProgress); err != nil {
		return err
	}

	// Get data from header and params
	var err error
	bookProgress.UserID, err = userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	bookProgress.LibraryID, bookProgress.BookID, err = getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	// Check permission
	if perm, err := permissionServ.Serv.GetUserLibraryPermission(bookProgress.UserID, bookProgress.LibraryID); err != nil {
		return err
	} else if !perm.CanReadBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to read books in this library")
	}

	updatedBookProgress, err := bookProgressServ.Serv.UpsertProgression(bookProgress)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(updatedBookProgress); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

// progressRead returns the user's progression on a book
func progressRead(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	// Check permission
	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanReadBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to read books in this library")
	}

	// RetrieveFromIDs
	progress, err := bookProgressServ.Serv.ReadProgressionFromIDs(userID, bookID, libraryID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(progress); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// CreateBookProgressHandlers create handlers for BookProgress
func CreateBookProgressHandlers(app *fiber.App) {
	bookProgressServ.NewService(database.BookProgressDB, database.BookDB)
	app.Get("/library/:library_id/book/:book_id/progress", userMiddleware.Protected(), progressRead)
	app.Post("/library/:library_id/book/:book_id/progress", userMiddleware.Protected(), progressUpsert)
}
