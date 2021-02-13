package handlers

import (
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/gofiber/fiber/v2"
)

// ProgressUpdate updates the user's progression on a book
func ProgressUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var progressData data.APIProgressData
	if err := ParseBodyJSON(ctx, &progressData); err != nil {
		return err
	}

	progressData.UserID = string(ctx.Request().Header.Peek("ID"))
	progressData.BookID = ctx.Params("book_id")
	progressData.LibraryID = ctx.Params("library_id")
	progressData.LastReadDate = time.Now()

	if progressData.Progress < 0 || progressData.Progress > 100 {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Progress is out of range: must be between 0 and 100")
	}

	if ok, err := internal.HasUserAccessToLibraryFromID(progressData.UserID, progressData.LibraryID); err != nil {
		return err
	} else if !ok {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User does not have access to the specified library.")
	}

	bookUserData, err := progressData.ToBookProgressData()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	userData, err := database.ProgressUpdateOrInsert(ctx.Context(), *bookUserData)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
