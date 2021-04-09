package handlers

import (
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database/book"
	grpcclient "github.com/alexandr-io/backend/media/grpc/client"
	"github.com/alexandr-io/backend/media/internal"

	"github.com/gofiber/fiber/v2"
)

// BookCover give the cover of a book
func BookCover(ctx *fiber.Ctx) error {
	bookData, err := book.GetFromID(ctx.Params("book_id"))
	if err != nil {
		return err
	}
	if bookData.CoverPath == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Cover does not exist")
	}

	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), string(ctx.Request().Header.Peek("ID")), bookData.LibraryID); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	file, err := internal.DownloadFile(ctx.Context(), path.Join(bookData.CoverPath))
	if err != nil {
		return err
	}

	err = ctx.Send(file.Data)
	ctx.Set(fiber.HeaderContentType, file.ContentType)
	//ctx.Set(fiber.HeaderContentDisposition, "attachment")

	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
