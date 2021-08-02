package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	userLibraryServ "github.com/alexandr-io/backend/library/internal/userlibrary"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// retrieveUserLibraries retrieve the list of libraries names the connect user has access to
func retrieveUserLibraries(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	result, err := userLibraryServ.Serv.ReadUserLibraryFromUserID(userID)
	if err != nil {
		return err
	}

	// Return the new library to the user
	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// CreateUserLibraryHandlers create handlers for UserLibrary
func CreateUserLibraryHandlers(app *fiber.App) {
	userLibraryServ.NewService(database.UserLibraryDB)
	app.Get("/libraries", userMiddleware.Protected(), retrieveUserLibraries)
}
