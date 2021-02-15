package library

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

// GetFromID retrieve a library from its ID
func GetFromID(libraryID string) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionLibrary)

	var DBLibrary data.Library

	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	libraryFilter := bson.D{{Key: "_id", Value: id}}
	if err := collection.FindOne(ctx, libraryFilter).Decode(&DBLibrary); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBLibrary, nil
}

// GetPermissionFromUserAndLibraryID find the user permissions for the given library and put it in the user.Permissions field
func GetPermissionFromUserAndLibraryID(user *data.User, libraryIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionLibraries)

	libraryID, err := primitive.ObjectIDFromHex(libraryIDStr)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	var libraryData data.UserLibrary
	userFilter := bson.D{{"user_id", userID}, {"library_id", libraryID}}
	result := collection.FindOne(ctx, userFilter)
	if result.Err() == mongo.ErrNoDocuments {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, result.Err().Error())
	}
	if err = result.Decode(&libraryData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	user.Permissions = libraryData.Permissions
	return nil
}
