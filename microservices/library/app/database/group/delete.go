package group

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Delete a group from it's ID
func Delete(groupID string) error {
	id, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	groupFilter := bson.D{{Key: "_id", Value: id}, {"priority", bson.D{{"$ne", -1}}}}
	var object permissions.Group
	if err := database.GroupCollection.FindOneAndDelete(context.Background(), groupFilter).Decode(&object); err != nil {
		if err == mongo.ErrNoDocuments {
			return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Group not found.")
		}
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	groupUpperFilter := bson.D{{"priority", bson.D{{"$gte", object.Priority}}}, {"library_id", object.LibraryID}}
	if _, err = database.GroupCollection.UpdateMany(context.Background(), groupUpperFilter, bson.D{{"$inc", bson.D{{"priority", -1}}}}); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
