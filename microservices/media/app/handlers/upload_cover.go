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

	"github.com/gofiber/fiber/v2"
)

// UploadBookCover upload the cover of a book from the request to the media folder and save path in database
func UploadBookCover(ctx *fiber.Ctx) error {
	// Parse the book data from the form
	bookData := data.Book{
		ID:        ctx.FormValue("book_id", ""),
		LibraryID: ctx.FormValue("library_id", ""),
	}

	// Check that data are not empty. Use same function than book upload
	if err := checkBookUploadBadInputs(bookData); err != nil {
		return err
	}

	// TODO: Upload authorization or other? Create gRPC route to retrieve all the permissions
	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), string(ctx.Request().Header.Peek("ID")), bookData.LibraryID); err != nil {
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

	// MEDIA_PATH/LibraryID/BookID_cover
	bookData.CoverPath = path.Join(os.Getenv("MEDIA_PATH"), bookData.LibraryID, bookData.ID+"_cover_"+cover.Filename)

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
	if _, err = book.Update(ctx.Context(), bookData); err != nil {
		if err := internal.DeleteFile(ctx.Context(), bookData.CoverPath); err != nil {
			return err
		}
		return err
	}

	// Send URL to retrieve image to library book metadata
	// http(s)://HOST/book/{bookID}/cover
	coverURL := string(ctx.Request().URI().Scheme()) + "://" + path.Join(string(ctx.Request().Host()), "book", bookData.ID, "cover")
	if err = grpcclient.CoverUploaded(ctx.Context(), bookData.ID, coverURL); err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
