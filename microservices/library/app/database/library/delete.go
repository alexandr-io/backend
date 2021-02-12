package library

import (
	"context"
	"errors"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete delete the library for a user and the name of the library.
// TODO: Remove unused field
func Delete(userID string, libraryID string) error {

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
		return errors.New("library does not exist")
	}

	collection = database.Instance.Db.Collection(database.CollectionLibraries)

	userLibraryFilter := bson.D{{"library_id", id}}
	_, err = collection.DeleteMany(ctx, userLibraryFilter)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
