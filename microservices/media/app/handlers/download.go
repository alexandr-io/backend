package handlers

import (
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database/book"
	grpcclient "github.com/alexandr-io/backend/media/grpc/client"
	"github.com/alexandr-io/backend/media/internal"

	"github.com/gofiber/fiber/v2"
)

// DownloadBook download a book from a book ID
func DownloadBook(ctx *fiber.Ctx) error {
	bookDB, err := book.GetFromID(ctx.Params("book_id"))
	if err != nil {
		return err
	}

	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), string(ctx.Request().Header.Peek("ID")), bookDB.LibraryID); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	file, err := internal.DownloadFile(ctx.Context(), path.Join(bookDB.Path))
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
