package user

import (
	"context"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete user delete a user corresponding to the given user id
func Delete(id primitive.ObjectID) error {
	result, err := database.UserCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: id},
		},
	)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "can't find user to delete")
	}
	return nil
}
