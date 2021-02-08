package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
)

// LibraryRetrieve retrieve the information of a library related to the connected user.
func LibraryRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryOwner := data.LibrariesOwner{
		UserID: userID,
	}

	libraryID := ctx.Params("library_id")

	library, err := database.GetLibraryByUserIDAndLibraryID(libraryOwner, libraryID)

	// library, err := database.GetLibraryByUserIDAndName(libraryOwner, *libraryName)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
