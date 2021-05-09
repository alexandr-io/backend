package invitation

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Delete delete an invitation corresponding to the given invitation token
func Delete(collection *mongo.Collection, token string) error {
	result, err := collection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "token", Value: token},
		},
	)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "can't find invitation to delete")
	}
	return nil
}
