package book

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetFromID retrieve a book from its ID
func GetFromID(bookID primitive.ObjectID) (*data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionBook)

	var bookData data.Book

	libraryFilter := bson.D{{Key: "_id", Value: bookID}}
	if err := collection.FindOne(ctx, libraryFilter).Decode(&bookData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Book not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &bookData, nil
}

// GetListFromLibraryID get the list of books in the given library
// TODO: pagination
func GetListFromLibraryID(libraryID primitive.ObjectID) (*[]data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionBook)

	var bookData []data.Book

	bookFilter := bson.D{{Key: "library_id", Value: libraryID}}
	cursor, err := collection.Find(ctx, bookFilter)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err = cursor.All(ctx, &bookData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &bookData, nil
}
