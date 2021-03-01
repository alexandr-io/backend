package userdata

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes the UserData entry in the database
func Delete(ctx context.Context, dataID string) error {
	collection := database.Instance.Db.Collection(database.CollectionUserData)

	id, err := primitive.ObjectIDFromHex(dataID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	if result, err := collection.DeleteOne(ctx, bson.D{{"_id", id}}); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found.")
	}

	return nil
}
