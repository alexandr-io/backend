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
	var object []data.Library

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	cursor, err := database.LibrariesCollection.Aggregate(context.Background(), mongo.Pipeline{
		bson.D{{"$match", bson.D{{"user_id", id}, {"invited_by", bson.D{{"$exists", false}}}}}},
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

	if err = cursor.All(context.Background(), &object); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Return the libraries object
	return &object, nil
}

// GetInvitedToFromUserID retrieve the user_libraries the current user is invited to.
func GetInvitedToFromUserID(userID string) (*[]data.Library, error) {
	var object []data.Library

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	cursor, err := database.LibrariesCollection.Aggregate(context.Background(), mongo.Pipeline{
		bson.D{{"$match", bson.D{{"user_id", id}, {"invited_by", bson.D{{"$exists", true}}}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "library"},
			{"localField", "library_id"},
			{"foreignField", "_id"},
			{"as", "library"},
		}}},
		bson.D{{"$unwind", bson.D{{"path", "$library"}}}},
		bson.D{{"$addFields", bson.D{{"library.invited_by", "$invited_by"}}}},
		bson.D{{"$replaceRoot", bson.D{{"newRoot", "$library"}}}},
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err = cursor.All(context.Background(), &object); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Return the libraries object
	return &object, nil
}

// GetFromUserIDAndLibraryID retrieve a user library from the user's ID and the library's ID
func GetFromUserIDAndLibraryID(userID string, libraryID string) (*data.UserLibrary, error) {
	var object data.UserLibrary

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	libraryObjID, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	filters := bson.D{{"user_id", userObjID}, {"library_id", libraryObjID}}
	if err := database.LibrariesCollection.FindOne(context.Background(), filters).Decode(&object); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return &object, nil
}
