package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
)

// BookRetrieve find and return the metadata of a book
func BookRetrieve(ctx *fiber.Ctx) error {
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

	if !user.CanSeeBooks() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to see the books in this library")
	}

	book, err := database.BookRetrieve(ctx.Context(), *bookData)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(book); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
