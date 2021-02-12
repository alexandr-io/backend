package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"
	"github.com/alexandr-io/backend/library/database/library"

	"github.com/gofiber/fiber/v2"
)

// BooksRetrieve retrieve the list of book in the library.
func BooksRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryID := ctx.Params("library_id")

	user := data.User{
		ID: userID,
	}
	err := library.GetPermissionFromUserAndLibraryID(&user, libraryID)
	if err != nil {
		return err
	}

	if !user.CanSeeBooks() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Unauthorised")
	}
	result, err := book.GetListFromLibraryID(libraryID)
	if err != nil {
		return err
	}

	// Return the new library to the user
	if err := ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
