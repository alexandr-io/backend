package handlers

import (
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getLibraryIDFromParams(ctx *fiber.Ctx) (primitive.ObjectID, error) {
	ID := ctx.Params("library_id")
	libraryID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return libraryID, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	return libraryID, nil
}

func getLibraryBookIDFromParams(ctx *fiber.Ctx) (libraryID primitive.ObjectID, bookID primitive.ObjectID, err error) {
	library := ctx.Params("library_id")
	book := ctx.Params("book_id")
	libraryID, err = primitive.ObjectIDFromHex(library)
	if err != nil {
		err = data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
		return
	}
	bookID, err = primitive.ObjectIDFromHex(book)
	if err != nil {
		err = data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
		return
	}
	return
}

func getGroupIDFromParams(ctx *fiber.Ctx) (primitive.ObjectID, error) {
	ID := ctx.Params("group_id")
	groupID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return groupID, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	return groupID, nil
}

func getDataIDFromParams(ctx *fiber.Ctx) (primitive.ObjectID, error) {
	ID := ctx.Params("data_id")
	dataID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return dataID, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	return dataID, nil
}
