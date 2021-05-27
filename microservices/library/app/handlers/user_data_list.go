package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/userdata"

	"github.com/gofiber/fiber/v2"
)

// UserDataList returns the user's data on a book
func UserDataList(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	userData, err := userdata.RetrieveManyFromIDs(userID, libraryID, bookID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
