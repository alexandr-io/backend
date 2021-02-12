package libraries

import (
	"context"
	"github.com/alexandr-io/backend/library/database"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

// GetFromUserID get the user_libraries the current user has access to.
func GetFromUserID(userID string) (*[]data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var object []data.Library

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	collection := mongo.Instance.Db.Collection(database.CollectionLibraries)
	cursor, err := collection.Aggregate(ctx, mongo2.Pipeline{
		bson.D{{"$match", bson.D{{"user_id", id}}}},
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

	err = cursor.All(ctx, &object)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	// Return the libraries object
	return &object, nil
}
