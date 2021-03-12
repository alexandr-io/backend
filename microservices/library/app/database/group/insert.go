package group

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert on the database a new book in a library.
func Insert(group permissions.Group) (*permissions.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionGroup)

	groupUpperFilter := bson.D{{"priority", bson.D{{"$gte", group.Priority}}}, {"library_id", group.LibraryID}}
	_, err := collection.UpdateMany(ctx, groupUpperFilter, bson.D{{"$inc", bson.D{{"priority", 1}}}})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	result, err := collection.InsertOne(ctx, group)
	if err != nil {
		return nil, err
	}

	group.ID = result.InsertedID.(primitive.ObjectID)
	return &group, nil
}
