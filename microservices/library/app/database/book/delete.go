package book

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/database/bookprogress"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete delete a book corresponding to the given invitation token
func Delete(bookID primitive.ObjectID) error {
	result, err := database.BookCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: bookID},
		},
	)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "can't find book to delete")
	}

	// TODO: create logic so that when book progress delete failed, the book previously deleted in restored
	if err = bookprogress.Delete(data.BookProgressData{
		BookID: bookID,
	}); err != nil {
		return err
	}
	return nil
}
