package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// BookDelete delete the metadata of a book
func BookDelete(ctx *fiber.Ctx) error {
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
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User does not have access to the specified library.")
	}

	err = database.BookDelete(ctx.Context(), *bookData)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
