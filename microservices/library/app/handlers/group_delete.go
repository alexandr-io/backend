package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/group"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// GroupDelete delete a group.
func GroupDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryID := ctx.Params("library_id")
	groupID := ctx.Params("group_id")

	if perm, err := internal.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if perm.CanManagePermissions() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to delete a group in this library")
	}

	err := group.Delete(groupID)
	if err != nil {
		return err
	}

	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
