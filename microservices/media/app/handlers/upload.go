package handlers

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database/book"
	grpcclient "github.com/alexandr-io/backend/media/grpc/client"
	"github.com/alexandr-io/backend/media/internal"
	"github.com/alexandr-io/backend/media/middleware"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadBook upload the book from the request to the media folder and create a database document
func UploadBook(ctx *fiber.Ctx) error {

	// Parse the bookDB data from the body
	id, err := primitive.ObjectIDFromHex(ctx.FormValue("book_id", ""))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	libraryID, err := primitive.ObjectIDFromHex(ctx.FormValue("library_id", ""))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	bookDB := data.Book{
		ID:        id,
		LibraryID: libraryID,
	}

	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), middleware.RetrieveAuthInfos(ctx).ID, bookDB.LibraryID); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	// Get the book from the context
	file, err := ctx.FormFile("book")
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	bookDB.Path = path.Join(os.Getenv("MEDIA_PATH"), bookDB.LibraryID.Hex(), bookDB.ID.Hex()+"_"+file.Filename)

	// Save file to /media directory:
	fd, err := file.Open()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	fileContent, err := ioutil.ReadAll(fd)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err = internal.UploadFile(ctx.Context(), fileContent, bookDB.Path); err != nil {
		return err
	}

	if _, err = book.Insert(bookDB); err != nil {
		if err = internal.DeleteFile(ctx.Context(), bookDB.Path); err != nil {
			return err
		}
		return err
	}

	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
