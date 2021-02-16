package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"
	"github.com/alexandr-io/backend/library/database/library"

	"github.com/gofiber/fiber/v2"
)

// BookUpdate update the metadata of a book
func BookUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryID := ctx.Params("library_id")
	bookIDStr := ctx.Params("book_id")

	var bookDB data.Book
	if err := ParseBodyJSON(ctx, &bookDB); err != nil {
		return err
	}

	var user = &data.User{ID: userID}
	if err := library.GetPermissionFromUserAndLibraryID(user, libraryID); err != nil {
		return err
	}

	if !user.CanUpdateBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to update books in this library")
	}
	bookDB.ID = bookIDStr

	bookData, err := bookDB.ToBookData()
	if err != nil {
		return err
	}
	result, err := book.Update(bookData)
	if err != nil {
		return err
	}
	if err := ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
