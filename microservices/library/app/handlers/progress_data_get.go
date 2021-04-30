package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/bookprogress"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// ProgressRetrieve returns the user's progression on a book
func ProgressRetrieve(ctx *fiber.Ctx) error {
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
	if perm, err := internal.GetUserLibraryPermission(userID.Hex(), libraryID.Hex()); err != nil {
		return err
	} else if perm.CanReadBook() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to read books in this library")
	}

	// RetrieveFromIDs
	progress, err := bookprogress.RetrieveFromIDs(userID, bookID, libraryID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(progress); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
