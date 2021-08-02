package handlers

import (
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	userDataServ "github.com/alexandr-io/backend/library/internal/user_data"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// parseUserDataRequest parses and handles error for creating a user_data
func parseUserDataRequest(ctx *fiber.Ctx, userData *data.UserData) error {
	if err := ParseBodyJSON(ctx, userData); err != nil {
		return err
	}
	if err := userData.ValidateUserDataType(); err != nil {
		return err
	}

	if userData.Type == "highlight" && userData.OffsetEnd == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Missing required parameter for highlights: offset_end")
	}

	return nil
}

// userDataCreate creates a UserData in the database.
func userDataCreate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	userDataRequest := data.UserData{
		UserID:           userID,
		LibraryID:        libraryID,
		BookID:           bookID,
		LastModifiedDate: time.Now(),
	}
	if err = parseUserDataRequest(ctx, &userDataRequest); err != nil {
		return err
	}

	userData, err := userDataServ.Serv.CreateUserData(userDataRequest)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusCreated).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

// userDataGet returns a user's data on a book
func userDataGet(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	_, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}
	dataID, err := getDataIDFromParams(ctx)
	if err != nil {
		return err
	}

	userData, err := userDataServ.Serv.ReadUserData(userID, bookID, dataID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// userDataList returns the user's data on a book
func userDataList(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	userData, err := userDataServ.Serv.ReadAllUserData(userID, libraryID, bookID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// userDataUpdate updates a UserData in the database
func userDataUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}
	dataID, err := getDataIDFromParams(ctx)
	if err != nil {
		return err
	}

	userDataRequest := data.UserData{
		UserID:           userID,
		LibraryID:        libraryID,
		BookID:           bookID,
		ID:               dataID,
		LastModifiedDate: time.Now(),
	}
	if err = parseUserDataRequest(ctx, &userDataRequest); err != nil {
		return err
	}

	userData, err := userDataServ.Serv.UpdateUserData(userDataRequest)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// userDataDeleteOne deletes a UserData from the database
func userDataDeleteOne(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}
	dataID, err := getDataIDFromParams(ctx)
	if err != nil {
		return err
	}

	err = userDataServ.Serv.DeleteUserData(userID, libraryID, bookID, dataID)
	if err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// userDataDeleteAllInBook deletes all UserData (from a specific book) from the database
func userDataDeleteAllInBook(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	err = userDataServ.Serv.DeleteAllUserDataIn(userID, libraryID, bookID)
	if err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// CreateUserDataHandlers create handlers for UserData
func CreateUserDataHandlers(app *fiber.App) {
	userDataServ.NewService(database.UserDataDB)
	app.Get("/library/:library_id/book/:book_id/data", userMiddleware.Protected(), userDataList)
	app.Post("/library/:library_id/book/:book_id/data", userMiddleware.Protected(), userDataCreate)
	app.Delete("/library/:library_id/book/:book_id/data", userMiddleware.Protected(), userDataDeleteAllInBook)
	app.Get("/library/:library_id/book/:book_id/data/:data_id", userMiddleware.Protected(), userDataGet)
	app.Post("/library/:library_id/book/:book_id/data/:data_id", userMiddleware.Protected(), userDataUpdate)
	app.Delete("/library/:library_id/book/:book_id/data/:data_id", userMiddleware.Protected(), userDataDeleteOne)
}
