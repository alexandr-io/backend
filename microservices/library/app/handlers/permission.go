package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	permissionServ "github.com/alexandr-io/backend/library/internal/permission"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// userLibraryPermissionRetrieve retrieve a user's permissions in a library.
func userLibraryPermissionRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	object, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(object); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// CreatePermissionHandlers create handlers for Permission
func CreatePermissionHandlers(app *fiber.App) {
	permissionServ.NewService(database.UserLibraryDB, database.GroupDB)
	app.Get("/library/:library_id/permissions", userMiddleware.Protected(), userLibraryPermissionRetrieve)
}
