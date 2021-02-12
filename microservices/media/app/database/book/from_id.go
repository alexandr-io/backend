package book

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"
	mongo2 "github.com/alexandr-io/backend/media/database/mongo"

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
	collection := mongo2.Instance.Db.Collection(database.CollectionBook)
	result := collection.FindOne(ctx, filter)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, result.Err().Error())
	}
	if err := result.Decode(&book); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &book, nil
}
