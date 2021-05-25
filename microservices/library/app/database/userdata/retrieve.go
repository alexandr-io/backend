package userdata

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

// RetrieveOneFromIDs retrieves the user's book data from the mongo database
func RetrieveOneFromIDs(userID, libraryID, bookID, dataID primitive.ObjectID) (*data.UserData, error) {
	/* TODO: Retrieve list? */

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionUserData)

	filter := bson.D{
		{"user_id", userID},
		{"book_id", bookID},
		{"library_id", libraryID},
		{"_id", dataID},
	}
	var result data.UserData
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &result, nil
}
