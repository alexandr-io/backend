package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LibraryInviteDecline decline an invitation to a library
func LibraryInviteDecline(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	if library, err := libraries.GetFromUserIDAndLibraryID(userID.Hex(), libraryID.Hex()); err != nil {
		return err
	} else if library.InvitedBy == primitive.NilObjectID {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "You already have access to this library")
	}

	if err := libraries.DeleteFromUserIDAndLibraryID(userID, libraryID); err != nil {
		return err
	}
	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
