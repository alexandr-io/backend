package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"
	"github.com/alexandr-io/backend/library/database/library"

	"github.com/gofiber/fiber/v2"
)

// BookRetrieve find and return the metadata of a book
func BookRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	bookData := &data.Book{
		ID:         ctx.Params("book_id"),
		LibraryID:  ctx.Params("library_id"),
		UploaderID: userID,
	}

	bookData.UploaderID = userID

	var user = &data.User{ID: userID}

	if err := library.GetPermissionFromUserAndLibraryID(user, bookData.LibraryID); err != nil {
		return err
	}

	if !user.CanSeeBooks() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to see the books in this library")
	}

	result, err := book.GetFromID(bookData.ID)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
