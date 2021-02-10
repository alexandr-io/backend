package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/database/library/getters"

	"github.com/gofiber/fiber/v2"
)

// LibraryRetrieve retrieve the information of a library related to the connected user.
func LibraryRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	libraryID := ctx.Params("library_id")
	user := &data.User{ID: string(ctx.Request().Header.Peek("ID"))}
	if err := database.GetLibraryPermission(user, &data.Library{
		ID: libraryID,
	}); err != nil {
		return err
	}
	library, err := getters.GetLibraryFromID(libraryID)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
