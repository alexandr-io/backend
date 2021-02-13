package handlers

import (
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

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

	user := data.User{ID: progressData.UserID}
	library := data.Library{ID: progressData.LibraryID}
	if err := database.GetLibraryPermission(&user, &library); err != nil {
		return err
	}
	if !user.CanReadBooks() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User cannot access this book")
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
