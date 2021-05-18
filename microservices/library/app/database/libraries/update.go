package libraries

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update a user's library.
func Update(library data.UserLibrary) (*data.UserLibrary, error) {
	filters := bson.D{{"user_id", library.UserID}, {"library_id", library.LibraryID}}
	if err := database.LibrariesCollection.FindOneAndUpdate(context.Background(), filters, bson.D{{"$set", library}}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&library); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User's library not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &library, nil
}
