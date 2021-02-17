package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"

	"github.com/gofiber/fiber/v2"
)

// LibraryDelete delete a library of the connected user.
func LibraryDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if err := library.Delete(ctx.Params("library_id")); err != nil {
		return err
	}

	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
