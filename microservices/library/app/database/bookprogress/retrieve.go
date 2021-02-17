package bookprogress

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Retrieve retrieves the user's book progress from the mongo database
func Retrieve(ctx context.Context, progressRetrieve data.APIProgressData) (*data.APIProgressData, error) {
	collection := database.Instance.Db.Collection(database.CollectionBookProgress)
	bookUserData, err := progressRetrieve.ToBookProgressData()

	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	filter := bson.D{
		{"user_id", bookUserData.UserID},
		{"book_id", bookUserData.BookID},
		{"library_id", bookUserData.LibraryID},
	}
	var result data.BookProgressData
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	progressData := result.ToAPIProgressData()
	return &progressData, nil
}
