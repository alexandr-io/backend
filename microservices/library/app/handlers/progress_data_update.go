package handlers

import (
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"
	"github.com/alexandr-io/backend/library/database/bookprogress"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// ProgressUpdate updates the user's progression on a book
func ProgressUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Parse data
	var progressData data.APIProgressData
	if err := ParseBodyJSON(ctx, &progressData); err != nil {
		return err
	}

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	// Fill data
	progressData.UserID = userID.Hex()
	progressData.BookID = bookID.Hex()
	progressData.LibraryID = libraryID.Hex()
	progressData.LastReadDate = time.Now()

	// Check permission
	if perm, err := internal.GetUserLibraryPermission(progressData.UserID, progressData.LibraryID); err != nil {
		return err
	} else if perm.CanReadBook() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to read books in this library")
	}

	bookUserData, err := progressData.ToBookProgressData()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	// Check book existence
	if _, err = book.GetFromID(bookID); err != nil {
		return err
	}

	// Update / Insert data
	userData, err := bookprogress.Upsert(ctx.Context(), *bookUserData)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
