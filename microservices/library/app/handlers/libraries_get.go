package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/gofiber/fiber/v2"
)

// LibrariesRetrieve retrieve the list of libraries names the connect user has access to.
func LibrariesRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	librariesOwner := data.LibrariesOwner{
		UserID: userID,
	}

	libraries, err := database.GetLibrariesNamesByUserID(librariesOwner)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Return the new library to the user
	if err := ctx.Status(fiber.StatusOK).JSON(libraries); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
