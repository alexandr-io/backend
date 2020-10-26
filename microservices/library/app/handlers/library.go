package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/gofiber/fiber/v2"
)

func LibraryRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	libraryOwner := new(data.LibraryOwner)
	if err := ParseBodyJSON(ctx, libraryOwner); err != nil {
		return err
	}

	library := data.LibraryOwner{
		Username: libraryOwner.Username,
	}

	// Return the new library to the user
	if err := ctx.Status(fiber.StatusOK).JSON(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

func LibraryCreation(ctx *fiber.Ctx) error {
	return nil
}
