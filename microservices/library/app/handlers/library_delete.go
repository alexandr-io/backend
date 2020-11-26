package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
)

// LibraryDelete delete a library of the connected user.
func LibraryDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	libraryName := new(data.LibraryName)

	if err := ParseBodyJSON(ctx, libraryName); err != nil {
		return err
	}
	libraryOwner := data.LibrariesOwner{
		UserID: userID,
	}
	err := database.DeleteLibrary(libraryOwner, *libraryName)
	if err != nil {
		return err
	}
	return nil
}
