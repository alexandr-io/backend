package database

import (
	"context"

	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserLibraryDB instance of UserLibraryCollection
var UserLibraryDB *UserLibraryCollection

// UserLibraryCollection contain the db collection for the user library interface
type UserLibraryCollection struct {
	collection *mongo.Collection
}

const userLibraryCollectionName = "user_library"

// NewUserLibraryCollection create a UserLibraryCollection
func NewUserLibraryCollection(db *mongo.Database) *UserLibraryCollection {
	return &UserLibraryCollection{collection: db.Collection(userLibraryCollectionName)}
}

// Create creates a document in the user_library collection
func (c *UserLibraryCollection) Create(userLibrary data.UserLibrary) (*data.UserLibrary, error) {
	insertedResult, err := c.collection.InsertOne(context.Background(), userLibrary)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	userLibrary.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &userLibrary, nil
}

// ReadFromUserID get the user libraries which the current user has access to
func (c *UserLibraryCollection) ReadFromUserID(userID string) (*[]data.Library, error) {
	var object []data.Library

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	cursor, err := c.collection.Aggregate(context.Background(), mongo.Pipeline{
		bson.D{{"$match", bson.D{{"user_id", id}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "library"},
			{"localField", "library_id"},
			{"foreignField", "_id"},
			{"as", "library"},
		}}},
		bson.D{{"$unwind", bson.D{{"path", "$library"}}}},
		bson.D{{"$replaceRoot", bson.D{{"newRoot", "$library"}}}},
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err = cursor.All(context.Background(), &object); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Return the libraries object
	return &object, nil
}

// ReadFromUserIDAndLibraryID retrieve a user library from the user's ID and the library's ID
func (c *UserLibraryCollection) ReadFromUserIDAndLibraryID(userID primitive.ObjectID, libraryID primitive.ObjectID) (*data.UserLibrary, error) {
	var object data.UserLibrary

	filters := bson.D{{"user_id", userID}, {"library_id", libraryID}}
	if err := c.collection.FindOne(context.Background(), filters).Decode(&object); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return &object, nil
}

// ReadFromLibraryID retrieve a user library from the library's ID
func (c *UserLibraryCollection) ReadFromLibraryID(libraryID primitive.ObjectID) (*data.UserLibrary, error) {
	var object data.UserLibrary

	filters := bson.D{{"library_id", libraryID}}
	if err := c.collection.FindOne(context.Background(), filters).Decode(&object); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return &object, nil
}

// Update a user library.
func (c *UserLibraryCollection) Update(library data.UserLibrary) (*data.UserLibrary, error) {
	filters := bson.D{{"user_id", library.UserID}, {"library_id", library.LibraryID}}
	if err := c.collection.FindOneAndUpdate(context.Background(), filters, bson.D{{"$set", library}}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&library); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User's library not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &library, nil
}

// Delete a user library
func (c *UserLibraryCollection) Delete(id primitive.ObjectID) error {
	userLibraryFilter := bson.D{{"library_id", id}}
	deleteResult, err := c.collection.DeleteOne(context.Background(), userLibraryFilter)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if deleteResult.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User's library not found.")
	}
	return nil
}
