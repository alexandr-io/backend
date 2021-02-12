package library

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/database/libraries"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert insert on the database a new library in the user's libraries.
func Insert(userIDStr string, DBLibrary data.Library) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionLibrary)

	insertedResult, err := collection.InsertOne(ctx, DBLibrary)
	if err != nil {
		return nil, err
	}
	DBLibrary.ID = insertedResult.InsertedID.(primitive.ObjectID).Hex()

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	userLibrary := data.UserLibrary{
		UserID:      userID,
		LibraryID:   insertedResult.InsertedID.(primitive.ObjectID),
		Permissions: []permissions.Permission{permissions.Owner},
	}
	_, err = libraries.Insert(userLibrary)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBLibrary, nil
}
