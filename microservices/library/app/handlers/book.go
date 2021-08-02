package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	bookServ "github.com/alexandr-io/backend/library/internal/book"
	permissionServ "github.com/alexandr-io/backend/library/internal/permission"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// bookCreation create the metadata of a book
func bookCreate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Parse data
	var book data.Book
	if err := ParseBodyJSON(ctx, &book); err != nil {
		return err
	}
	if book.Title == "" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Book title is required")
	}

	// Get data from header and params
	var err error
	book.UploaderID, err = userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	book.LibraryID, err = getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(book.UploaderID, book.LibraryID); err != nil {
		return err
	} else if !perm.CanUploadBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to upload books on this library")
	}

	result, err := bookServ.Serv.CreateBook(book)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusCreated).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// bookRead find and return the metadata of a book
func bookRead(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanDisplayBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to see the books in this library")
	}

	// Retrieve
	result, err := bookServ.Serv.ReadFromID(bookID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// booksRead retrieve the list of book in a library
func booksRead(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanDisplayBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to see the books in this library")
	}

	result, err := bookServ.Serv.ReadFromLibraryID(libraryID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// bookUpdate update the metadata of a book
func bookUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	// Parse data
	var book data.Book
	if err = ParseBodyJSON(ctx, &book); err != nil {
		return err
	}
	book.ID = bookID

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanUpdateBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to see the books in this library")
	}

	// Update
	result, err := bookServ.Serv.UpdateBook(book)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// bookDelete delete the metadata of a book
func bookDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, bookID, err := getLibraryBookIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanDeleteBook() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to see the books in this library")
	}

	// Delete
	if err = bookServ.Serv.DeleteBook(bookID); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// CreateBookHandlers create handlers for Book
func CreateBookHandlers(app *fiber.App) {
	bookServ.NewService(database.BookDB, database.BookProgressDB)
	app.Post("/library/:library_id/book", userMiddleware.Protected(), bookCreate)
	app.Get("/library/:library_id/book/:book_id", userMiddleware.Protected(), bookRead)
	app.Get("/library/:library_id/books", userMiddleware.Protected(), booksRead)
	app.Post("/library/:library_id/book/:book_id", userMiddleware.Protected(), bookUpdate)
	app.Delete("/library/:library_id/book/:book_id", userMiddleware.Protected(), bookDelete)
}
