package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/gofiber/fiber/v2"
)

// BookCreation create the metadata of a book
func BookCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	bookData := new(data.BookCreation)
	if err := ParseBodyJSON(ctx, bookData); err != nil {
		return err
	}

	bookData.UploaderID = userID
	bookData.LibraryID = ctx.Params("library_id")

	var user = &data.User{ID: userID}
	var library = &data.Library{ID: bookData.LibraryID}
	err := database.GetLibraryPermission(user, library)
	if err != nil {
		return err
	}

	if !user.CanUploadBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to upload books on this library")
	}

	book, err := database.BookCreate(ctx.Context(), *bookData)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusCreated).JSON(book); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
