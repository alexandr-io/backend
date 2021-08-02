package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	libraryServ "github.com/alexandr-io/backend/library/internal/library"
	permissionServ "github.com/alexandr-io/backend/library/internal/permission"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// libraryCreate create a library for the connected user
func libraryCreate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	var library data.Library
	if err = ParseBodyJSON(ctx, &library); err != nil {
		return err
	}

	response, err := libraryServ.Serv.CreateLibrary(library, userID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusCreated).JSON(response); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// libraryRead retrieve the information of a library related to the connected user
func libraryRead(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	// Check permission
	if _, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	}

	// Retrieve
	result, err := libraryServ.Serv.ReadLibrary(libraryID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// libraryDelete delete a library of the connected user
func libraryDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanDeleteLibrary() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to delete this library")
	}

	if err = libraryServ.Serv.DeleteLibrary(libraryID); err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

func CreateLibraryHandlers(app *fiber.App) {
	libraryServ.NewService(database.LibraryDB, database.UserLibraryDB, database.GroupDB, database.BookProgressDB)
	app.Post("/library", userMiddleware.Protected(), libraryCreate)
	app.Get("/library/:library_id", userMiddleware.Protected(), libraryRead)
	app.Delete("/library/:library_id", userMiddleware.Protected(), libraryDelete)
}
