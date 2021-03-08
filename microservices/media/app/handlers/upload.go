package handlers

import (
	grpcclient "github.com/alexandr-io/backend/media/grpc/client"
	"io/ioutil"
	"os"
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database/book"
	"github.com/alexandr-io/backend/media/internal"
	"github.com/gofiber/fiber/v2"
)

// UploadBook upload the book from the request to the media folder and create a database document
func UploadBook(ctx *fiber.Ctx) error {

	// Parse the bookDB data from the body
	bookDB := data.Book{
		ID:        ctx.FormValue("book_id", ""),
		LibraryID: ctx.FormValue("library_id", ""),
	}

	if isAllowed, err := grpcclient.UploadAuthorization(ctx.Context(), string(ctx.Request().Header.Peek("ID")), bookDB.LibraryID); err != nil {
		return err
	} else if !isAllowed {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Not authorized")
	}

	if err := checkBookUploadBadInputs(bookDB); err != nil {
		return err
	}

	// Get the book from the context
	file, err := ctx.FormFile("book")
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	bookDB.Path = path.Join(os.Getenv("MEDIA_PATH"), bookDB.LibraryID, bookDB.ID+"_"+file.Filename)

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
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
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

func checkBookUploadBadInputs(book data.Book) error {
	errorList := make(map[string]string)

	if book.ID == "" {
		errorList["book_id"] = "The field is required"
	}

	if book.LibraryID == "" {
		errorList["library_id"] = "The field is required"
	}

	if len(errorList) != 0 {
		errorInfo := data.NewErrorInfo(string(badInputsJSON(errorList)), 0)
		errorInfo.ContentType = fiber.MIMEApplicationJSON
		return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
	}
	return nil
}
