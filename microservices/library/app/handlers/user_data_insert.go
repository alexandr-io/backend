package handlers

import (
	"strings"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/userdata"

	"github.com/gofiber/fiber/v2"
)

func validateUserDataType(userType string) error {
	for _, dataType := range data.UserDataTypes {
		if userType == dataType {
			return nil
		}
	}
	return data.NewHTTPErrorInfo(fiber.StatusBadRequest,
		"type parameter must be one of: "+strings.Join(data.UserDataTypes[:], ", "))
}

// parseUserDataRequest parses and handles error for creating a user_data
func parseUserDataRequest(ctx *fiber.Ctx, userData *data.UserData) error {
	if err := ParseBodyJSON(ctx, userData); err != nil {
		return err
	}
	if err := validateUserDataType(userData.Type); err != nil {
		return err
	}

	if userData.Type == "highlight" && userData.OffsetEnd == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing required parameter for highlights: offset_end")
	}

	return nil
}

// UserDataCreate creates a UserData in the database.
func UserDataCreate(ctx *fiber.Ctx) error {
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

	now := time.Now()
	userDataRequest := data.UserData{
		UserID:           userID,
		LibraryID:        libraryID,
		BookID:           bookID,
		LastModifiedDate: now,
	}
	if err := parseUserDataRequest(ctx, &userDataRequest); err != nil {
		return err
	}

	userData, err := userdata.Insert(userDataRequest)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusCreated).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
