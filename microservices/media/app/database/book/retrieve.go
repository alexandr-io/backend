package book

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetFromID retrieve a book from the ID given by the library MS.
func GetFromID(bookID string) (*data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book data.Book
	filter := bson.D{{Key: "book_id", Value: bookID}}
	collection := database.Instance.Db.Collection(database.CollectionBook)

	if err := collection.FindOne(ctx, filter).Decode(&book); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Book not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &book, nil
}
