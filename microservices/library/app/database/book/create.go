package book

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert insert on the database a new book in a library.
func Insert(bookData data.Book) (*data.Book, error) {
	result, err := database.BookCollection.InsertOne(context.Background(), bookData)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	bookData.ID = result.InsertedID.(primitive.ObjectID)
	return &bookData, nil
}
