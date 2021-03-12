package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/group"
	"github.com/alexandr-io/backend/library/database/libraries"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// GroupJoin add a user to a group.
func GroupJoin(ctx *fiber.Ctx) error {

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryID := ctx.Params("library_id")
	groupID := ctx.Params("group_id")

	if perm, err := internal.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if perm.CanManagePermissions() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to add a user in a group in this library")
	}

	var groupJoinData data.GroupJoinData
	if err := ParseBodyJSON(ctx, &groupJoinData); err != nil {
		return err
	}

	if _, err := libraries.GetFromUserIDAndLibraryID(groupJoinData.UserID, libraryID); err != nil {
		return err
	}

	if _, err := group.AddUserToGroup(groupJoinData.UserID, groupID, libraryID); err != nil {
		return err
	}
	return nil
}
