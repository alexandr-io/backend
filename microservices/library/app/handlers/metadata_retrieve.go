package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	grpcclient "github.com/alexandr-io/backend/library/grpc/client"

	"github.com/gofiber/fiber/v2"
)

// MetadataRetrieve find and return the metadata of a book on Google Books API
func MetadataRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	title := ctx.Query("title")
	authors := ctx.Query("authors")
	response, err := grpcclient.Metadata(title, authors)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err := ctx.Status(fiber.StatusOK).JSON(response); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
