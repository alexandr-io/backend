package handlers

import (
	"errors"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/gofiber/fiber/v2"
)

// swagger:route GET /libraries LIBRARY libraries_list
// Retrieve the libraries names of the connected user
// responses:
//	200: librariesNamesResponse
//	400: badRequestErrorResponse

// LibrariesRetrieve retrieve the list of libraries names the connect user has access to.
func LibrariesRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	librariesOwner := new(data.LibrariesOwner)

	librariesOwner.UserID = userID

	libraries, err := database.GetLibrariesNamesByUserID(*librariesOwner)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Return the new library to the user
	if err := ctx.Status(fiber.StatusOK).JSON(libraries); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// swagger:route GET /library LIBRARY library_retrieve
// Retrieve a library from it's name
// responses:
//	200: libraryResponse
//	400: badRequestErrorResponse

// LibraryRetrieve retrieve the information of a library related to the connected user.
func LibraryRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryOwner := &data.LibrariesOwner{
		UserID: userID,
	}
	libraryName := new(data.LibraryName)

	if err := ParseBodyJSON(ctx, libraryName); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	library, err := database.GetLibraryByUserIDAndName(*libraryOwner, *libraryName)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err := ctx.Status(fiber.StatusOK).JSON(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// swagger:route POST /library LIBRARY library_creation
// Create a library for the connected user
// responses:
//	200: librariesNamesResponse
//	400: badRequestErrorResponse

// LibraryCreation create a library for the connected user
func LibraryCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	library := new(data.Library)

	if err := ParseBodyJSON(ctx, library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	library.Books = []data.BookInfo{}

	libraryOwner := &data.LibrariesOwner{
		UserID: userID,
	}

	_, err := database.InsertLibrary(*libraryOwner, *library)
	if err != nil {
		var badInputError *data.BadInputError
		if errors.As(err, &badInputError) {
			errorInfo := data.NewErrorInfo(string(badInputError.JSONError), 0)
			errorInfo.ContentType = fiber.MIMEApplicationJSON
			return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
		}
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	libraryName := &data.LibraryName{
		Name: library.Name,
	}
	library, err = database.GetLibraryByUserIDAndName(*libraryOwner, *libraryName)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err := ctx.Status(fiber.StatusOK).JSON(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// LibraryDelete delete a library of the connected user.
func LibraryDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	libraryName := new(data.LibraryName)

	if err := ParseBodyJSON(ctx, libraryName); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	libraryOwner := &data.LibrariesOwner{
		UserID: userID,
	}
	err := database.DeleteLibrary(*libraryOwner, *libraryName)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
