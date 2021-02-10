package getters

import (
	"context"
	"log"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/alexandr-io/backend/library/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetLibraryFromID retrieve a library from its ID
func GetLibraryFromID(libraryID string) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(library.Collection)

	var DBLibrary data.Library

	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}

	log.Println(id)
	libraryFilter := bson.D{{Key: "_id", Value: id}}
	if err := collection.FindOne(ctx, libraryFilter).Decode(&DBLibrary); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBLibrary, nil
}
