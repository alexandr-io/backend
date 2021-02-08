package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"

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

	if ok, err := internal.CanUserModifyBook(userID, libraryID, bookIDStr); err != nil {
		return err
	} else if !ok {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You do not have access to this library.")
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
