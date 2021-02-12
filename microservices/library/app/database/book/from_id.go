package book

import (
	"context"
	"github.com/alexandr-io/backend/library/database"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

// GetFromID retrieve a book from its ID
func GetFromID(bookID string) (*data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(database.CollectionBook)

	var DBBook data.Book

	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	libraryFilter := bson.D{{Key: "_id", Value: id}}
	result := collection.FindOne(ctx, libraryFilter)
	if result.Err() == mongo2.ErrNoDocuments {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, result.Err().Error())
	}
	if err := result.Decode(&DBBook); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBBook, nil
}
