package book

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert create a book on the database
func Insert(book data.Book) (*data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bookCollection := database.Instance.Db.Collection(database.CollectionBook)
	insertedResult, err := bookCollection.InsertOne(ctx, book)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	book.ID = insertedResult.InsertedID.(primitive.ObjectID).Hex()
	return &book, nil
}
