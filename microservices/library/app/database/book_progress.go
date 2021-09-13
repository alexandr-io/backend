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

// BookProgressDB instance of BookProgressCollection
var BookProgressDB *BookProgressCollection

// BookProgressCollection contain the db collection for the book progress interface
type BookProgressCollection struct {
	collection *mongo.Collection
}

const bookProgressCollectionName = "book_progress"

// NewBookProgressCollection create a BookProgressCollection
func NewBookProgressCollection(db *mongo.Database) *BookProgressCollection {
	return &BookProgressCollection{collection: db.Collection(bookProgressCollectionName)}
}

// Upsert updates the user's book progress in the database
// If no document is found in the library, a new document is created
func (c *BookProgressCollection) Upsert(bookProgress data.BookProgressData) (*data.BookProgressData, error) {
	filter := bson.D{
		{"user_id", bookProgress.UserID},
		{"book_id", bookProgress.BookID},
		{"library_id", bookProgress.LibraryID},
	}

	if err := c.collection.FindOneAndUpdate(context.Background(), filter, bson.D{{"$set", bookProgress}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1),
	).Decode(&bookProgress); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &bookProgress, nil
}

// ReadFromBookID retrieves the user's book progress from the mongo database
func (c *BookProgressCollection) ReadFromBookID(bookID primitive.ObjectID) (*data.BookProgressData, error) {
	filter := bson.D{
		{"book_id", bookID},
	}

	var result data.BookProgressData
	if err := c.collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &result, nil
}

// ReadFromLibraryID retrieves the user's book progress from the mongo database
func (c *BookProgressCollection) ReadFromLibraryID(libraryID primitive.ObjectID) (*data.BookProgressData, error) {
	filter := bson.D{
		{"library_id", libraryID},
	}

	var result data.BookProgressData
	if err := c.collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &result, nil
}

// ReadFromIDs retrieves the user's book progress from the mongo database
func (c *BookProgressCollection) ReadFromIDs(userID primitive.ObjectID, bookID primitive.ObjectID, libraryID primitive.ObjectID) (*data.BookProgressData, error) {
	filter := bson.D{
		{"user_id", userID},
		{"book_id", bookID},
		{"library_id", libraryID},
	}

	var result data.BookProgressData
	if err := c.collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &result, nil
}

// Delete deletes the user's book progress data entry in the database
func (c *BookProgressCollection) Delete(bookProgress data.BookProgressData) error {
	if result, err := c.collection.DeleteOne(context.Background(), bookProgress); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Progress not found.")
	}

	return nil
}
