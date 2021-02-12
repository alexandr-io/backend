package libraries

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert create a document in the user_library collection
func Insert(DBLibrary data.UserLibrary) (*data.UserLibrary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionLibraries)
	insertedResult, err := collection.InsertOne(ctx, DBLibrary)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	DBLibrary.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &DBLibrary, nil
}
