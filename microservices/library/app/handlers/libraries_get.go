package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"

	"github.com/gofiber/fiber/v2"
)

// LibrariesRetrieve retrieve the list of libraries names the connect user has access to.
func LibrariesRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	libraryHasAccess, err := libraries.GetFromUserID(userID)
	if err != nil {
		return err
	}
	libraryIsInvited, err := libraries.GetInvitedToFromUserID(userID)
	if err != nil {
		return err
	}

	result := data.Libraries{
		HasAccess: libraryHasAccess,
		IsInvited: libraryIsInvited,
	}

	// Return the new library to the user
	if err := ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
