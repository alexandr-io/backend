package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/gofiber/fiber/v2"
)

// BookUpdate update the metadata of a book
func BookUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID := string(ctx.Request().Header.Peek("ID"))
	libraryID := ctx.Params("library_id")
	bookIDStr := ctx.Params("book_id")

	book := new(data.Book)
	if err := ParseBodyJSON(ctx, book); err != nil {
		return err
	}

	var user = &data.User{ID: userID}
	var library = &data.Library{ID: libraryID}
	err := database.GetLibraryPermission(user, library)
	if err != nil {
		return err
	}

	if !user.CanUpdateBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to update books in this library")
	}
	book.ID = bookIDStr

	bookResult, err := database.BookUpdate(ctx.Context(), libraryID, *book)
	if err != nil {
		return err
	}
	if err := ctx.Status(fiber.StatusOK).JSON(bookResult); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
