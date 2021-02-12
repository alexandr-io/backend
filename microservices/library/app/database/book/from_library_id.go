package book

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetListFromLibraryID get the list of books in the given library
// TODO: pagination
func GetListFromLibraryID(libraryID string) (*[]data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionBook)

	var DBBook []data.Book

	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	bookFilter := bson.D{{Key: "library_id", Value: id}}
	cursor, err := collection.Find(ctx, bookFilter)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err := cursor.All(ctx, &DBBook); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBBook, nil
}
