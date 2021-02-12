package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/gofiber/fiber/v2"
)

// BookCreation create the metadata of a book
func BookCreation(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))

	var bookDB data.Book
	if err := ParseBodyJSON(ctx, &bookDB); err != nil {
		return err
	}

	bookDB.UploaderID = userID
	bookDB.LibraryID = ctx.Params("library_id")

	var user = &data.User{ID: userID}
	err := library.GetPermissionFromUserAndLibraryID(user, bookDB.LibraryID)
	if err != nil {
		return err
	}

	if !user.CanUploadBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to upload books on this library")
	}

	bookData, err := bookDB.ToBookData()
	if err != nil {
		return err
	}

	result, err := book.Insert(bookData)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusCreated).JSON(result.ToBook()); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
