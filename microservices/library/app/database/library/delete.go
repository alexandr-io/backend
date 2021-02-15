package library

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/database/bookprogress"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete delete the library for a user and the name of the library.
func Delete(libraryID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionLibrary)

	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	libraryFilter := bson.D{{Key: "_id", Value: id}}
	deleteResult, err := collection.DeleteOne(ctx, libraryFilter)
	if err != nil {
		return err
	} else if deleteResult.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library not found.")
	}

	collection = database.Instance.Db.Collection(database.CollectionLibraries)

	userLibraryFilter := bson.D{{"library_id", id}}
	deleteResult, err = collection.DeleteMany(ctx, userLibraryFilter)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if deleteResult.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User's library not found.")
	}

	_ = bookprogress.Delete(ctx, data.BookProgressData{
		LibraryID: id,
	})
	return nil
}
