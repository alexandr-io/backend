package library

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/database/bookprogress"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete delete the library for a user and the name of the library.
func Delete(libraryID string) error {
	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	libraryFilter := bson.D{{Key: "_id", Value: id}}
	deleteResult, err := database.LibraryCollection.DeleteOne(context.Background(), libraryFilter)
	if err != nil {
		return err
	} else if deleteResult.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library not found.")
	}

	userLibraryFilter := bson.D{{"library_id", id}}
	deleteResult, err = database.LibrariesCollection.DeleteMany(context.Background(), userLibraryFilter)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if deleteResult.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User's library not found.")
	}

	_ = bookprogress.Delete(data.BookProgressData{
		LibraryID: id,
	})
	return nil
}
