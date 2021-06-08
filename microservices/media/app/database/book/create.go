package book

import (
	"context"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"

	"github.com/gofiber/fiber/v2"
)

// Insert create a book on the database
func Insert(book data.Book) (*data.Book, error) {
	_, err := database.BookCollection.InsertOne(context.Background(), book)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &book, nil
}
