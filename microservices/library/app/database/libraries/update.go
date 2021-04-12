package libraries

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update a user's library.
func Update(library data.UserLibrary) (*data.UserLibrary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionLibraries)

	filters := bson.D{{"user_id", library.UserID}, {"library_id", library.LibraryID}}
	if err := collection.FindOneAndUpdate(ctx, filters, bson.D{{"$set", library}}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&library); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User's library not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &library, nil
}
