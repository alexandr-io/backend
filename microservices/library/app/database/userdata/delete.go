package userdata

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes the UserData entry in the database
func Delete(ctx context.Context, userID, libraryID, bookID, dataID primitive.ObjectID) error {
	collection := database.Instance.Db.Collection(database.CollectionUserData)

	filter := bson.D{
		{"user_id", userID},
		{"library_id", libraryID},
		{"book_id", bookID},
		{"_id", dataID},
	}

	if result, err := collection.DeleteOne(ctx, filter); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found.")
	}

	return nil
}

// DeleteAllIn deletes all UserData entry (in a book/library) in the database
func DeleteAllIn(ctx context.Context, userID, libraryID, bookID primitive.ObjectID) error {
	collection := database.Instance.Db.Collection(database.CollectionUserData)

	filter := bson.D{
		{"user_id", userID},
		{"library_id", libraryID},
		{"book_id", bookID},
	}

	if result, err := collection.DeleteMany(ctx, filter); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found.")
	}

	return nil
}
