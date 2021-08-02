package database

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var LibraryDB *LibraryCollection

type LibraryCollection struct {
	collection *mongo.Collection
}

const libraryCollectionName = "library"

func NewLibraryCollection(db *mongo.Database) *LibraryCollection {
	return &LibraryCollection{collection: db.Collection(libraryCollectionName)}
}

// Create insert a new library in the user's libraries.
func (c *LibraryCollection) Create(library data.Library) (*data.Library, error) {
	insertedResult, err := c.collection.InsertOne(context.Background(), library)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	library.ID = insertedResult.InsertedID.(primitive.ObjectID).Hex()
	return &library, nil
}

// Read retrieve a library from its ID
func (c *LibraryCollection) Read(libraryID primitive.ObjectID) (*data.Library, error) {
	var library data.Library

	libraryFilter := bson.D{{Key: "_id", Value: libraryID}}
	if err := c.collection.FindOne(context.Background(), libraryFilter).Decode(&library); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &library, nil
}

// Delete the library for a user and the name of the library.
func (c *LibraryCollection) Delete(libraryID primitive.ObjectID) error {
	libraryFilter := bson.D{{Key: "_id", Value: libraryID}}
	deleteResult, err := c.collection.DeleteOne(context.Background(), libraryFilter)
	if err != nil {
		return err
	} else if deleteResult.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library not found.")
	}
	return nil
}
