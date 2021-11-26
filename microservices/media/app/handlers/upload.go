package handlers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database/book"
	grpcclient "github.com/alexandr-io/backend/media/grpc/client"
	"github.com/alexandr-io/backend/media/internal"
	"github.com/alexandr-io/backend/media/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/h2non/filetype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadBook upload the book from the request to the media folder and create a database document
func UploadBook(ctx *fiber.Ctx) error {

	// Parse the bookData data from the body
	id, err := primitive.ObjectIDFromHex(ctx.FormValue("book_id", ""))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	libraryID, err := primitive.ObjectIDFromHex(ctx.FormValue("library_id", ""))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	bookData := data.Book{
		ID:        id,
		LibraryID: libraryID,
	}

	// check permission
	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), middleware.RetrieveAuthInfos(ctx).ID, bookData.LibraryID); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	// Get the book from the context
	file, err := ctx.FormFile("book")
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	bookData.Path = path.Join(os.Getenv("MEDIA_PATH"), bookData.LibraryID.Hex(), bookData.ID.Hex()+"_"+file.Filename)

	// open and read file
	fd, err := file.Open()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	fileContent, err := ioutil.ReadAll(fd)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Get file type and check type
	kind, _ := filetype.Match(fileContent)
	fileType := ""
	if kind.MIME.Value == "application/pdf" {
		fileType = "pdf"
	} else if (kind.MIME.Value == "application/epub+zip" || kind.MIME.Value == "application/zip") && filepath.Ext(file.Filename) == ".epub" {
		fileType = "epub"
	} else {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, fmt.Sprintf("file type %s %s not supported", kind.Extension, kind.MIME.Value))
	}

	// Save file to /media directory:
	if err = internal.UploadFile(ctx.Context(), fileContent, bookData.Path); err != nil {
		return err
	}

	// Insert in DB
	if _, err = book.Insert(bookData); err != nil {
		if err = internal.DeleteFile(ctx.Context(), bookData.Path); err != nil {
			return err
		}
		return err
	}

	// Send type to library MS
	if err = grpcclient.BookUploaded(ctx.Context(), bookData.ID, fileType); err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
