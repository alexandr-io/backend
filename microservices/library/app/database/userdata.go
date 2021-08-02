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

var UserDataDB *UserDataCollection

type UserDataCollection struct {
	collection *mongo.Collection
}

const userDataCollectionName = "user_data"

func NewUserDataCollection(db *mongo.Database) *UserDataCollection {
	return &UserDataCollection{collection: db.Collection(userDataCollectionName)}
}

// Create insert a new user's book data to the mongo database
func (c *UserDataCollection) Create(userData data.UserData) (*data.UserData, error) {
	insertedResult, err := c.collection.InsertOne(context.Background(), userData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	userData.ID = insertedResult.InsertedID.(primitive.ObjectID)

	return &userData, nil
}

// Read retrieves the user's book data from the mongo database
func (c *UserDataCollection) Read(userID, bookID, dataID primitive.ObjectID) (*data.UserData, error) {
	filter := bson.D{
		{"user_id", userID},
		{"book_id", bookID},
		{"_id", dataID},
	}
	var result data.UserData
	if err := c.collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	result.ID = dataID

	return &result, nil
}

// ReadAll retrieves the user's book data from the mongo database
func (c *UserDataCollection) ReadAll(userID, libraryID, bookID primitive.ObjectID) (*[]data.UserData, error) {
	filter := bson.D{
		{"user_id", userID},
		{"book_id", bookID},
		{"library_id", libraryID},
	}
	var result []data.UserData
	cursor, err := c.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &result, nil
}

// Update updates the user's book progress in the database
func (c *UserDataCollection) Update(userData data.UserData) (*data.UserData, error) {
	filter := bson.D{
		{"user_id", userData.UserID},
		{"book_id", userData.BookID},
		{"library_id", userData.LibraryID},
		{"_id", userData.ID},
	}

	if err := c.collection.FindOneAndUpdate(context.Background(), filter, bson.D{{"$set", userData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&userData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, err.Error())
	}

	return &userData, nil
}

// Delete deletes the UserData entry in the database
func (c *UserDataCollection) Delete(userID, libraryID, bookID, dataID primitive.ObjectID) error {
	filter := bson.D{
		{"user_id", userID},
		{"library_id", libraryID},
		{"book_id", bookID},
		{"_id", dataID},
	}

	if result, err := c.collection.DeleteOne(context.Background(), filter); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found.")
	}

	return nil
}

// DeleteAllIn deletes all UserData entry (in a book/library) in the database
func (c *UserDataCollection) DeleteAllIn(userID, libraryID, bookID primitive.ObjectID) error {
	filter := bson.D{
		{"user_id", userID},
		{"library_id", libraryID},
		{"book_id", bookID},
	}

	if result, err := c.collection.DeleteMany(context.Background(), filter); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found.")
	}

	return nil
}
