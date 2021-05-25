package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/userdata"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDataUpdate updates a UserData in the database.
func UserDataUpdate(ctx *fiber.Ctx) error {
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
	dataID, err := primitive.ObjectIDFromHex(ctx.Params("data_id"))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	userDataRequest := data.UserData{
		UserID:    userID,
		LibraryID: libraryID,
		BookID:    bookID,
		ID:        dataID,
	}
	if err := parseUserDataRequest(ctx, &userDataRequest); err != nil {
		return err
	}

	userData, err := userdata.Update(ctx.Context(), userDataRequest)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
