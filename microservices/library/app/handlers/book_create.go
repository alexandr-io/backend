package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// BookCreation create the metadata of a book
func BookCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Parse data
	var bookData data.Book
	if err := ParseBodyJSON(ctx, &bookData); err != nil {
		return err
	}
	if bookData.Title == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Book title is required")
	}

	// Get data from header and params
	var err error
	bookData.UploaderID, err = userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	bookData.LibraryID, err = getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := internal.GetUserLibraryPermission(bookData.UploaderID.Hex(), bookData.LibraryID.Hex()); err != nil {
		return err
	} else if perm.CanUploadBook() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to upload books on this library")
	}

	result, err := book.Insert(bookData)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusCreated).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
