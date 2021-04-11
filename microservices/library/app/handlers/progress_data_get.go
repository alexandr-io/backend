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

	// Retrieve
	progressRetrieve := data.APIProgressData{ // TODO create only one function with custom marshal like book example: https://github.com/awslabs/goformation/blob/master/cloudformation/custom_resource.go
		UserID:    userID.Hex(),
		BookID:    bookID.Hex(),
		LibraryID: libraryID.Hex(),
	}
	progress, err := bookprogress.Retrieve(ctx.Context(), progressRetrieve)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(progress); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
