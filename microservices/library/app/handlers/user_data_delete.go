package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// UserDataDelete deletes a UserData from the database.
func UserDataDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	//userData := data.APIUserData{
	//	UserID:    string(ctx.Request().Header.Peek("ID")),
	//	LibraryID: ctx.Params("library_id"),
	//	BookID:    ctx.Params("book_id"),
	//	ID:        ctx.Params("data_id"),
	//}

	return nil
}