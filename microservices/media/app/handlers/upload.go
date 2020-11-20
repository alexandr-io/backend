package handlers

import (
	"github.com/alexandr-io/backend/media/database"
	"github.com/alexandr-io/backend/media/internal"
	"io/ioutil"
	"os"
	"path"

	"github.com/alexandr-io/backend/media/data"
	"github.com/gofiber/fiber/v2"
)

func UploadBook(ctx *fiber.Ctx) error {

	// Parse the book data from the body
	book := new(data.Book)
	book.ID = ctx.FormValue("book_id", "")
	book.LibraryID = ctx.FormValue("library_id", "")

	if book.ID == "" || book.LibraryID == "" {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "Missing field: book_id or library_id")
	}

	// Get the book in from the context
	file, err := ctx.FormFile("book")
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())

	}

	book.Path = path.Join(os.Getenv("MEDIA_PATH"), book.LibraryID, book.ID+"_"+file.Filename)

	// Save file to /media directory:
	fd, err := file.Open()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	fileContent, err := ioutil.ReadAll(fd)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	err = internal.UploadFile(ctx.Context(), fileContent, book.Path)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	err = database.InsertBook(ctx.Context(), *book)
	if err != nil {
		err2 := internal.DeleteFile(ctx.Context(), book.Path)
		if err2 != nil {
			return err2
		}
		return err
	}

	return nil
}
