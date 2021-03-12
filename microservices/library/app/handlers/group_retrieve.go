package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database/group"

	"github.com/gofiber/fiber/v2"
)

// GroupRetrieve retrieve a group.
func GroupRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	libraryID := ctx.Params("library_id")
	groupID := ctx.Params("group_id")

	var groupData permissions.Group
	if err := ParseBodyJSON(ctx, &groupData); err != nil {
		return err
	}

	object, err := group.GetFromIDAndLibraryID(groupID, libraryID)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(object); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
