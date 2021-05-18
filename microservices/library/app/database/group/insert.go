package group

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert on the database a new book in a library.
func Insert(group permissions.Group) (*permissions.Group, error) {
	groupUpperFilter := bson.D{{"priority", bson.D{{"$gte", group.Priority}}}, {"library_id", group.LibraryID}}
	if _, err := database.GroupCollection.UpdateMany(context.Background(), groupUpperFilter, bson.D{{"$inc", bson.D{{"priority", 1}}}}); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	result, err := database.GroupCollection.InsertOne(context.Background(), group)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	group.ID = result.InsertedID.(primitive.ObjectID)
	return &group, nil
}
