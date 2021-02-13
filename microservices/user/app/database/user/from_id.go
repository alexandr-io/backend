package user

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FromID get an user by it's given id.
func FromID(userID string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionUser)

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	filter := bson.D{{Key: "_id", Value: id}}
	var result data.User

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
