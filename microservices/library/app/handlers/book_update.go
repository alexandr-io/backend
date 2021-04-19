package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// BookUpdate update the metadata of a book
func BookUpdate(ctx *fiber.Ctx) error {
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

	// Parse data
	var bookData data.Book
	if err = ParseBodyJSON(ctx, &bookData); err != nil {
		return err
	}
	bookData.ID = bookID

	// Check permission
	if perm, err := internal.GetUserLibraryPermission(userID.Hex(), libraryID.Hex()); err != nil {
		return err
	} else if perm.CanDeleteBook() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to update books in this library")
	}

	// Update
	result, err := book.Update(bookData)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
