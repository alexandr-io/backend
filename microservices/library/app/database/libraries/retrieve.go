package libraries

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetFromUserID get the user_libraries the current user has access to.
func GetFromUserID(userID string) (*[]data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var object []data.Library

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	collection := database.Instance.Db.Collection(database.CollectionLibraries)
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{
		bson.D{{"$match", bson.D{{"user_id", id}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "library"},
			{"localField", "library_id"},
			{"foreignField", "_id"},
			{"as", "library"},
		}}},
		bson.D{{"$unwind", bson.D{{"path", "$library"}}}},
		bson.D{{"$replaceRoot", bson.D{{"newRoot", "$library"}}}},
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err = cursor.All(ctx, &object); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Return the libraries object
	return &object, nil
}

// GetFromUserIDAndLibraryID retrieve a user library from the user's ID and the library's ID
func GetFromUserIDAndLibraryID(userID string, libraryID string) (*data.UserLibrary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var object data.UserLibrary

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	libraryObjID, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	collection := database.Instance.Db.Collection(database.CollectionLibraries)

	filters := bson.D{{"user_id", userObjID}, {"library_id", libraryObjID}}
	if err := collection.FindOne(ctx, filters).Decode(&object); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return &object, nil
}
