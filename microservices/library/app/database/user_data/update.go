package user_data

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update updates the user's book progress in the database
// If no document is found in the library, a new document is created
func Update(ctx context.Context, userData data.UserData) (*data.UserData, error) {
	collection := database.Instance.Db.Collection(database.CollectionUserData)

	filter := bson.D{
		{"_id", userData.ID},
		//{"user_id", userData.UserID},
		//{"book_id", userData.BookID},
		//{"library_id", userData.LibraryID},
	}

	if err := collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", userData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&userData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &userData, nil
}
