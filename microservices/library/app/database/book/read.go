package book

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetFromID retrieve a book from its ID
func GetFromID(bookID primitive.ObjectID) (*data.Book, error) {
	var bookData data.Book

	libraryFilter := bson.D{{Key: "_id", Value: bookID}}
	if err := database.BookCollection.FindOne(context.Background(), libraryFilter).Decode(&bookData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Book not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &bookData, nil
}

// GetBooksFromLibraryID get the list of books in the given library
// TODO: pagination
// TODO: add not found error
func GetBooksFromLibraryID(libraryID primitive.ObjectID) (*[]data.Book, error) {
	var bookData []data.Book

	bookFilter := bson.D{{Key: "library_id", Value: libraryID}}
	cursor, err := database.BookCollection.Find(context.Background(), bookFilter)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err = cursor.All(context.Background(), &bookData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &bookData, nil
}
