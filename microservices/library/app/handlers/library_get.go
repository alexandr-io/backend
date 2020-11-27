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
	libraryName := new(data.LibraryName)

	if err := ParseBodyJSON(ctx, libraryName); err != nil {
		return err
	}

	library, err := database.GetLibraryByUserIDAndName(libraryOwner, *libraryName)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
