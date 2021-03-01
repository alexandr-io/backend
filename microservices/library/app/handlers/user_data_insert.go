package handlers

import (
	"strings"

	"github.com/alexandr-io/backend/library/data"
	"github.com/gofiber/fiber/v2"
)

// parseRequest parses and handles error for creating a user_data
func parseRequest(ctx *fiber.Ctx, userData *data.APIUserData) error {
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
	if userData.Offset == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing required parameter: offset")
	}
	if userData.Type == "highlight" && userData.OffsetEnd == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing required parameter for highlights: offset_end")
	}

	return nil
}

// UserDataInsert creates a UserData in the database.
func UserDataInsert(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userData := data.APIUserData{
		UserID:    string(ctx.Request().Header.Peek("ID")),
		LibraryID: ctx.Params("library_id"),
		BookID:    ctx.Params("book_id"),
	}

	if err := parseRequest(ctx, &userData); err != nil {
		return err
	}
	return nil
}
