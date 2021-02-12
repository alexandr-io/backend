package library

import (
	"context"
	"github.com/alexandr-io/backend/library/database"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetFromID retrieve a library from its ID
func GetFromID(libraryID string) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(database.CollectionLibrary)

	var DBLibrary data.Library

	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}

	libraryFilter := bson.D{{Key: "_id", Value: id}}
	if err := collection.FindOne(ctx, libraryFilter).Decode(&DBLibrary); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBLibrary, nil
}
