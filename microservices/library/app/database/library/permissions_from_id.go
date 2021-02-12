package library

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	mongo2 "github.com/alexandr-io/backend/library/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPermissionFromUserAndLibraryID find the user permissions for the given library and put it in the user.Permissions field
func GetPermissionFromUserAndLibraryID(user *data.User, libraryIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo2.Instance.Db.Collection(database.CollectionLibraries)

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
