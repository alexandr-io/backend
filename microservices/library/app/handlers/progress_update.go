package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// ProgressUpdate updates the user's progression on a book
func ProgressUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	progressData := new(data.APIProgressData)
	if err := ParseBodyJSON(ctx, progressData); err != nil {
		return err
	}

	progressData.UserID = userID

	if progressData.Progress < 0 || progressData.Progress > 1 {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Progress is out of range: must be between 0 and 1")
	}

	ok, err := internal.HasUserAccessToLibraryFromID(userID, progressData.LibraryID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if !ok {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User does not have access to the specified library.")
	}

	userData, err := database.ProgressUpdate(ctx.Context(), *progressData)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
