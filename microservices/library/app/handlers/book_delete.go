package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
)

// BookDelete delete the metadata of a book
func BookDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	bookData := &data.BookRetrieve{
		ID:         ctx.Params("book_id"),
		LibraryID:  ctx.Params("library_id"),
		UploaderID: userID,
	}

	bookData.UploaderID = userID

	var user = &data.User{ID: userID}
	var library = &data.Library{ID: bookData.LibraryID}
	err := database.GetLibraryPermission(user, library)
	if err != nil {
		return err
	}

	if !user.CanDeleteBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to delete books on this library")
	}

	if err := database.BookDelete(ctx.Context(), *bookData); err != nil {
		return err
	}
	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
