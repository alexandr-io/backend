package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/group"
	"github.com/alexandr-io/backend/library/database/libraries"

	"github.com/gofiber/fiber/v2"
)

// GroupsRetrieveUser retrieve a list of the user's groups.
func GroupsRetrieveUser(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryID := ctx.Params("library_id")

	userLibraries, err := libraries.GetFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return err
	}

	object, err := group.GetFromIDListAndLibraryID(userLibraries.Groups, libraryID)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(object); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
