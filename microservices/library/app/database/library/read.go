package library

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetFromID retrieve a library from its ID
func GetFromID(libraryID primitive.ObjectID) (*data.Library, error) {
	var DBLibrary data.Library

	libraryFilter := bson.D{{Key: "_id", Value: libraryID}}
	if err := database.LibraryCollection.FindOne(context.Background(), libraryFilter).Decode(&DBLibrary); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBLibrary, nil
}
