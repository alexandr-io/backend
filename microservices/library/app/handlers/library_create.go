package handlers

import (
	"errors"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/gofiber/fiber/v2"
)

// LibraryCreation create a library for the connected user
func LibraryCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	var libraryDB data.Library
	if err := ParseBodyJSON(ctx, &libraryDB); err != nil {
		return err
	}

	response, err := library.Insert(userID, libraryDB)
	if err != nil {
		var badInputError *data.BadInputError
		if errors.As(err, &badInputError) {
			errorInfo := data.NewErrorInfo(string(badInputError.JSONError), 0)
			errorInfo.ContentType = fiber.MIMEApplicationJSON
			return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
		}
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err := ctx.Status(fiber.StatusCreated).JSON(response); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
