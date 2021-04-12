package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database/group"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GroupUpdate update a group.
func GroupUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryID := ctx.Params("library_id")
	groupID := ctx.Params("group_id")

	if perm, err := internal.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if perm.CanManagePermissions() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to add a user in a group in this library")
	}

	var groupData permissions.Group
	if err := ParseBodyJSON(ctx, &groupData); err != nil {
		return err
	}

	libraryObjID, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	groupData.LibraryID = libraryObjID

	groupObjID, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	groupData.ID = groupObjID

	object, err := group.Update(groupData)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(object); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
