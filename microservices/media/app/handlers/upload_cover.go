package handlers

import (
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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadBookCover upload the cover of a book from the request to the media folder and save path in database
func UploadBookCover(ctx *fiber.Ctx) error {
	// Parse the book data from the form
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

	// TODO: Upload authorization or other? Create gRPC route to retrieve all the permissions
	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), middleware.RetrieveAuthInfos(ctx).ID, bookData.LibraryID); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	cover, err := ctx.FormFile("cover")
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	if filepath.Ext(cover.Filename) != ".jpg" &&
		filepath.Ext(cover.Filename) != ".jpeg" &&
		filepath.Ext(cover.Filename) != ".png" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "cover format should be jpg, jpeg or png")
	}

	// MEDIA_PATH/{libraryID}/{bookID}_cover_{filename}
	bookData.CoverPath = path.Join(os.Getenv("MEDIA_PATH"), bookData.LibraryID.Hex(), bookData.ID.Hex()+"_cover_"+cover.Filename)

	// open and read cover file
	fd, err := cover.Open()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	coverContent, err := ioutil.ReadAll(fd)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Upload cover
	if err = internal.UploadFile(ctx.Context(), coverContent, bookData.CoverPath); err != nil {
		return err
	}

	// Add cover path in DB
	if _, err = book.Update(bookData); err != nil {
		if err := internal.DeleteFile(ctx.Context(), bookData.CoverPath); err != nil {
			return err
		}
		return err
	}

	// Send URL to retrieve image to library book metadata
	// http(s)://HOST/media/{bookData.CoverPath}
	scheme := "http"
	if os.Getenv("DEV") == "" {
		scheme = scheme + "s"
	}
	coverURL := scheme + "://" + path.Join(string(ctx.Request().Host()), bookData.CoverPath)
	if err = grpcclient.CoverUploaded(ctx.Context(), bookData.ID, coverURL); err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
