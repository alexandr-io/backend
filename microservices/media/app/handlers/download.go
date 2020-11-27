package handlers

import (
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"
	"github.com/alexandr-io/backend/media/internal"

	"github.com/gofiber/fiber/v2"
)

// DownloadBook download a book from a book ID
func DownloadBook(ctx *fiber.Ctx) error {

	book := new(data.Book)
	if err := ParseBodyJSON(ctx, book); err != nil {
		return err
	}

	book, err := database.GetBookByID(book)
	if err != nil {
		return err
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
