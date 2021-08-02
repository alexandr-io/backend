package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	metadataServ "github.com/alexandr-io/backend/library/internal/metadata"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// metadataRetrieve find and return the metadata of a book on Google Books API
func metadataRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	title := ctx.Query("title")
	authors := ctx.Query("authors")

	response, err := metadataServ.Serv.RequestMetadata(title, authors)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(response); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

func CreateMetadataHandlers(app *fiber.App) {
	metadataServ.NewService()
	app.Get("/metadata", userMiddleware.Protected(), metadataRetrieve)
}
