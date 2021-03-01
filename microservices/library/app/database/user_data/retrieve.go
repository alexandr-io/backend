package user_data

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Retrieve retrieves the user's book progress from the mongo database
func Retrieve(ctx context.Context, dataID string) (*data.APIUserData, error) {
	collection := database.Instance.Db.Collection(database.CollectionUserData)

	id, err := primitive.ObjectIDFromHex(dataID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	var result data.UserData
	filter := bson.D{{"_id", id}}
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	progressData := result.ToAPIUserData()
	return &progressData, nil
}

// TODO: Retrieve list
