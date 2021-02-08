package handlers

import (
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"
	"github.com/alexandr-io/backend/media/internal"
	"github.com/alexandr-io/backend/media/kafka/producers"

	"github.com/gofiber/fiber/v2"
)

// DownloadBook download a book from a book ID
func DownloadBook(ctx *fiber.Ctx) error {

	bookID := ctx.Params("book_id")
	book := &data.Book{
		ID: bookID,
	}
	if err := ParseBodyJSON(ctx, book); err != nil {
		return err
	}

	book, err := database.GetBookByID(book)
	if err != nil {
		return err
	}

	if isAllowed, err := producers.LibraryUploadAuthorizationRequestHandler(book, string(ctx.Request().Header.Peek("ID"))); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	file, err := internal.DownloadFile(ctx.Context(), path.Join(book.Path))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	err = ctx.Send(file.Data)
	ctx.Set(fiber.HeaderContentType, file.ContentType)
	ctx.Set(fiber.HeaderContentDisposition, "attachment")

	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
