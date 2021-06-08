package book

import (
	"context"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetFromID retrieve a book from the ID given by the library MS.
func GetFromID(bookID primitive.ObjectID) (*data.Book, error) {
	var book data.Book
	filter := bson.D{{Key: "book_id", Value: bookID}}

	if err := database.BookCollection.FindOne(context.Background(), filter).Decode(&book); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Book not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &book, nil
}
