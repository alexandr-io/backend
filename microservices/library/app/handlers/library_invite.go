package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"
	"github.com/alexandr-io/backend/library/grpc/client"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LibraryInvite invite a user to a library
func LibraryInvite(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}
	if perm, err := internal.GetUserLibraryPermission(userID.Hex(), libraryID.Hex()); err != nil {
		return err
	} else if perm.CanInviteUser() == false {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to invite user in this library")
	}

	var libraryInvite data.LibraryInvite
	if err := ParseBodyJSON(ctx, &libraryInvite); err != nil {
		return err
	}

	user, err := client.UserFromLogin(ctx.Context(), libraryInvite.Login)
	if err != nil {
		return err
	}

	InvitedUserID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	userLibrary, err := libraries.Insert(data.UserLibrary{
		UserID:      InvitedUserID,
		LibraryID:   libraryID,
		Permissions: libraryInvite.Permissions,
		Groups:      libraryInvite.Groups,
		InvitedBy:   userID,
	})
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusCreated).JSON(userLibrary); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
