package bookprogress

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Upsert updates the user's book progress in the database
// If no document is found in the library, a new document is created
func Upsert(ctx context.Context, bookUserData data.BookProgressData) (*data.BookProgressData, error) {
	collection := database.Instance.Db.Collection(database.CollectionBookProgress)

	filter := bson.D{
		{"user_id", bookUserData.UserID},
		{"book_id", bookUserData.BookID},
		{"library_id", bookUserData.LibraryID},
	}

	if err := collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", bookUserData}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1),
	).Decode(&bookUserData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &bookUserData, nil
}
