package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// BookCreation create the metadata of a book
func BookCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	book := new(data.BookCreation)
	if err := ParseBodyJSON(ctx, book); err != nil {
		return err
	}
	book.LibraryID = ctx.Params("library_id")

	book.UploaderID = userID

	ok, err := internal.HasUserAccessToLibraryFromID(userID, book.LibraryID)
	if err != nil {
		return err
	}
	if !ok {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User does not have access to the specified library.")
	}

	bookData, err := database.BookCreate(ctx.Context(), *book)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusCreated).JSON(bookData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
