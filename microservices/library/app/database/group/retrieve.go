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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetFromIDAndLibraryID retrieve a group from its ID and the library ID
func GetFromIDAndLibraryID(groupID string, libraryID string) (*permissions.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionGroup)

	var group permissions.Group

	groupObjID, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	libraryObjID, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	libraryFilter := bson.D{{Key: "_id", Value: groupObjID}, {"library_id", libraryObjID}}
	if err := collection.FindOne(ctx, libraryFilter).Decode(&group); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Group not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &group, nil
}

// GetFromIDListAndLibraryID retrieve a list of groups from a list of group IDs and the library ID
func GetFromIDListAndLibraryID(groupIDs []primitive.ObjectID, libraryID string) (*[]permissions.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionGroup)

	var group []permissions.Group

	libraryObjID, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	var filter bson.D
	if len(groupIDs) > 0 {
		filter = bson.D{{"$or", []interface{}{bson.D{{"priority", 0}}, bson.D{{"_id", bson.D{{"$in", groupIDs}}}, {"library_id", libraryObjID}}}}}
	} else {
		filter = bson.D{{"priority", 0}, {"library_id", libraryObjID}}
	}

	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(bson.D{{"priority", -1}}))
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err := cursor.All(ctx, &group); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &group, nil
}
