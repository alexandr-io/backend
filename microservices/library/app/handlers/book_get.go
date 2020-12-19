package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"
	"github.com/gofiber/fiber/v2"
)

// BookRetrieve find and return the metadata of a book
func BookRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	bookData := new(data.BookRetrieve)
	if err := ParseBodyJSON(ctx, bookData); err != nil {
		return err
	}

	bookData.UploaderID = userID

	ok, err := internal.HasUserAccessToLibraryFromID(userID, bookData.LibraryID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if !ok {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Invalid or expired JWT")
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
