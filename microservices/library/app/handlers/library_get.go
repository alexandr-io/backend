package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/gofiber/fiber/v2"
)

// LibraryRetrieve retrieve the information of a library related to the connected user.
func LibraryRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	libraryID := ctx.Params("library_id")
	user := &data.User{ID: string(ctx.Request().Header.Peek("ID"))}
	if err := library.GetPermissionFromUserAndLibraryID(user, libraryID); err != nil {
		return err
	}
	result, err := library.GetFromID(libraryID)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
