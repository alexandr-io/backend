package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/bookprogress"
	"github.com/alexandr-io/backend/library/database/library"

	"github.com/gofiber/fiber/v2"
)

// ProgressRetrieve returns the user's progression on a book
func ProgressRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	progressRetrieve := data.APIProgressData{
		UserID:    string(ctx.Request().Header.Peek("ID")),
		BookID:    ctx.Params("book_id"),
		LibraryID: ctx.Params("library_id"),
	}

	if progressRetrieve.BookID == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing mandatory parameter: book_id")
	}
	if progressRetrieve.LibraryID == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing mandatory parameter: library_id")
	}

	user := data.User{ID: progressRetrieve.UserID}

	if err := library.GetPermissionFromUserAndLibraryID(&user, progressRetrieve.LibraryID); err != nil {
		return err
	}
	if !user.CanReadBooks() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User cannot access this book")
	}

	progress, err := bookprogress.Retrieve(ctx.Context(), progressRetrieve)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(progress); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
