package library

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetFromID retrieve a library from its ID
func GetFromID(libraryID primitive.ObjectID) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionLibrary)

	var DBLibrary data.Library

	libraryFilter := bson.D{{Key: "_id", Value: libraryID}}
	if err := collection.FindOne(ctx, libraryFilter).Decode(&DBLibrary); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBLibrary, nil
}
