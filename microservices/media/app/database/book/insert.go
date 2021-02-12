package book

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"
	mongo2 "github.com/alexandr-io/backend/media/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert create a book on the database
func Insert(book data.Book) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bookCollection := mongo2.Instance.Db.Collection(database.CollectionBook)
	insertedResult, err := bookCollection.InsertOne(ctx, book)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return insertedResult, nil
}
