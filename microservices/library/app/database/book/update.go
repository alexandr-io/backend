package book

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update update a book in a library.
func Update(bookData data.Book) (*data.Book, error) {
	if err := database.BookCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", bookData.ID},
		},
		bson.D{{"$set", bookData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&bookData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Book not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &bookData, nil
}
