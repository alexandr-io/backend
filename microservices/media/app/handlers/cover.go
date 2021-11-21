package handlers

import (
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database/book"
	grpcclient "github.com/alexandr-io/backend/media/grpc/client"
	"github.com/alexandr-io/backend/media/internal"
	"github.com/alexandr-io/backend/media/middleware"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookCover give the cover of a book
func BookCover(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("book_id"))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	bookData, err := book.GetFromID(id)
	if err != nil {
		return err
	}
	if bookData.CoverPath == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Cover does not exist")
	}

	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), middleware.RetrieveAuthInfos(ctx).ID, bookData.LibraryID); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	file, err := internal.DownloadFile(ctx.Context(), path.Join(bookData.CoverPath))
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).Send(file.Data); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	ctx.Set(fiber.HeaderContentType, file.ContentType)
	return nil
}
