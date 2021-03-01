package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/alexandr-io/backend/library/database/user_data"

	"github.com/gofiber/fiber/v2"
)

// UserDataGet returns the user's data on a book
func UserDataGet(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userDataRetrieve := data.APIUserData{
		UserID:    string(ctx.Request().Header.Peek("ID")),
		LibraryID: ctx.Params("library_id"),
		BookID:    ctx.Params("book_id"),
		ID:        ctx.Params("data_id"),
	}

	if userDataRetrieve.ID == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing mandatory parameter: data_id")
	}
	if userDataRetrieve.BookID == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing mandatory parameter: book_id")
	}
	if userDataRetrieve.LibraryID == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing mandatory parameter: library_id")
	}

	user := data.User{ID: userDataRetrieve.UserID}

	if err := library.GetPermissionFromUserAndLibraryID(&user, userDataRetrieve.LibraryID); err != nil {
		return err
	}
	if !user.CanReadBooks() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "User cannot access this book")
	}

	userData, err := user_data.Retrieve(ctx.Context(), userDataRetrieve.ID)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
