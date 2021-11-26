package libraries

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteFromUserIDAndLibraryID delete a user library from a user ID and a library ID
func DeleteFromUserIDAndLibraryID(userID primitive.ObjectID, libraryID primitive.ObjectID) error {
	collection := database.Instance.Db.Collection(database.CollectionLibraries)

	result, err := collection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "user_id", Value: userID},
			{Key: "library_id", Value: libraryID},
		},
	)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "can't find library to delete")
	}

	return nil
}
