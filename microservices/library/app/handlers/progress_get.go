package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// ProgressRetrieve returns the user's progression on a book
func ProgressRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	progressRetrieve := new(data.APIProgressRetrieve)
	if err := ParseBodyJSON(ctx, progressRetrieve); err != nil {
		return err
	}

	progressRetrieve.UserID = userID

	ok, err := internal.HasUserAccessToLibraryFromID(progressRetrieve.UserID, progressRetrieve.LibraryID)
	if err != nil {
		return err
	}
	if !ok {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User does not have access to the specified library.")
	}

	progress, err := database.ProgressRetrieve(ctx.Context(), *progressRetrieve)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(progress); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
