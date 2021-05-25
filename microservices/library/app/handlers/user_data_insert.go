package handlers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"

	"github.com/alexandr-io/backend/library/data"
	"github.com/gofiber/fiber/v2"
)

// parseRequest parses and handles error for creating a user_data
func parseRequest(ctx *fiber.Ctx, userData *data.UserData) error {
	if err := ParseBodyJSON(ctx, &userData); err != nil {
		return err
	}

	for _, dataType := range data.UserDataTypes {
		if userData.Type == dataType {
			break
		}
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest,
			"type parameter must be one of: "+strings.Join(data.UserDataTypes[:], ", "))
	}

	if userData.Name == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing required parameter: name")
	}
	if userData.Offset == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing required parameter: offset")
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
	_, err = primitive.ObjectIDFromHex(ctx.Params("data_id"))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	userData := data.UserData{
		UserID:    userID,
		LibraryID: libraryID,
		BookID:    bookID,
	}

	if err := parseRequest(ctx, &userData); err != nil {
		return err
	}
	return nil
}
