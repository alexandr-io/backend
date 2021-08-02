package database

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var GroupDB *GroupCollection

type GroupCollection struct {
	collection *mongo.Collection
}

const groupCollectionName = "group"

func NewGroupCollection(db *mongo.Database) *GroupCollection {
	return &GroupCollection{collection: db.Collection(groupCollectionName)}
}

// Create a new book in a library
func (c *GroupCollection) Create(group permissions.Group) (*permissions.Group, error) {
	// Increase priority of existing groups of library if > group.Priority
	groupUpperFilter := bson.D{
		{"priority", bson.D{
			{"$gte", group.Priority},
		}},
		{"library_id", group.LibraryID},
	}
	if _, err := c.collection.UpdateMany(
		context.Background(),
		groupUpperFilter,
		bson.D{{"$inc", bson.D{{"priority", 1}}}},
	); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// insert new group
	result, err := c.collection.InsertOne(context.Background(), group)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	group.ID = result.InsertedID.(primitive.ObjectID)
	return &group, nil
}

// ReadFromIDListAndLibraryID retrieve a list of groups from a list of group IDs and the library ID
func (c *GroupCollection) ReadFromIDListAndLibraryID(groupIDs []primitive.ObjectID, libraryID primitive.ObjectID) (*[]permissions.Group, error) {
	var group []permissions.Group

	var filter bson.D
	if len(groupIDs) > 0 {
		filter = bson.D{{"library_id", libraryID}, {"$or", []interface{}{bson.D{{"priority", -1}}, bson.D{{"_id", bson.D{{"$in", groupIDs}}}}}}}
	} else {
		filter = bson.D{{"priority", 0}, {"library_id", libraryID}}
	}

	cursor, err := c.collection.Find(context.Background(), filter, options.Find().SetSort(bson.D{{"priority", -1}}))
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err = cursor.All(context.Background(), &group); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &group, nil
}

// ReadFromIDAndLibraryID retrieve a group from its ID and the library ID
func (c *GroupCollection) ReadFromIDAndLibraryID(groupID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.Group, error) {
	var group permissions.Group

	libraryFilter := bson.D{{Key: "_id", Value: groupID}, {"library_id", libraryID}}
	if err := c.collection.FindOne(context.Background(), libraryFilter).Decode(&group); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Group not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &group, nil
}

// Update a group
func (c *GroupCollection) Update(group permissions.Group) (*permissions.Group, error) {
	filters := bson.D{{"_id", group.ID}}
	if err := c.collection.FindOneAndUpdate(context.Background(), filters, bson.D{{"$set", group}}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&group); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Group not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &group, nil
}

// Delete a group
func (c *GroupCollection) Delete(groupID primitive.ObjectID) error {
	groupFilter := bson.D{{Key: "_id", Value: groupID}, {"priority", bson.D{{"$ne", -1}}}}
	var object permissions.Group
	if err := c.collection.FindOneAndDelete(context.Background(), groupFilter).Decode(&object); err != nil {
		if err == mongo.ErrNoDocuments {
			return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Group not found.")
		}
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// decrease priority of other groups
	groupUpperFilter := bson.D{{"priority", bson.D{{"$gte", object.Priority}}}, {"library_id", object.LibraryID}}
	if _, err := c.collection.UpdateMany(context.Background(), groupUpperFilter, bson.D{{"$inc", bson.D{{"priority", -1}}}}); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
