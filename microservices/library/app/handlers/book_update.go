package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/internal"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

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

	bookID, err := primitive.ObjectIDFromHex(bookIDStr)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	book.ID = bookID

	if err := database.BookUpdate(ctx.Context(), libraryID, *book); err != nil {
		return err
	}
	return nil
}
