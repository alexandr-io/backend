package userdata

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert creates a UserData entry in the database
func Insert(ctx context.Context, userData data.UserData) (*data.UserData, error) {
	collection := database.Instance.Db.Collection(database.CollectionUserData)

	result, err := collection.InsertOne(ctx, userData)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	userData.ID = result.InsertedID.(primitive.ObjectID)
	return &userData, nil
}
