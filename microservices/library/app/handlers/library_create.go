package handlers

import (
	"errors"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
)

// LibraryCreation create a library for the connected user
func LibraryCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	library := new(data.Library)

	if err := ParseBodyJSON(ctx, library); err != nil {
		return err
	}
	library.Books = []data.Book{}

	libraryOwner := data.LibrariesOwner{
		UserID: userID,
	}

	_, err := database.InsertLibrary(libraryOwner, *library)
	if err != nil {
		var badInputError *data.BadInputError
		if errors.As(err, &badInputError) {
			errorInfo := data.NewErrorInfo(string(badInputError.JSONError), 0)
			errorInfo.ContentType = fiber.MIMEApplicationJSON
			return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
		}
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	libraryName := &data.LibraryName{
		Name: library.Name,
	}
	library, err = database.GetLibraryByUserIDAndName(libraryOwner, *libraryName)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err := ctx.Status(fiber.StatusCreated).JSON(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
