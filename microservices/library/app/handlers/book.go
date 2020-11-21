package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"
	"github.com/gofiber/fiber/v2"
)

// BookRetrieve find anr return the metadata of a book
func BookRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	bookData := new(data.BookRetrieve)
	if err := ParseBodyJSON(ctx, bookData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
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

// BookCreation create the metadata of a book
func BookCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	bookData := new(data.BookCreation)
	if err := ParseBodyJSON(ctx, bookData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	bookData.UploaderID = userID

	ok, err := internal.HasUserAccessToLibraryFromID(userID, bookData.LibraryID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if !ok {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User does not have access to the specified library.")
	}

	book, err := database.BookCreate(ctx.Context(), *bookData)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(book); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// BookDelete delete the metadata of a book
func BookDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	bookData := new(data.BookRetrieve)
	if err := ParseBodyJSON(ctx, bookData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
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
		return err
	}
	return nil
}