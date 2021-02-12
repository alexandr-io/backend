package libraries

import (
	"context"
	"github.com/alexandr-io/backend/library/database"
	"time"

	"github.com/alexandr-io/backend/library/data"
	mongo2 "github.com/alexandr-io/backend/library/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(DBLibrary data.UserLibrary) (*data.UserLibrary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo2.Instance.Db.Collection(database.CollectionLibraries)
	insertedResult, err := collection.InsertOne(ctx, DBLibrary)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	DBLibrary.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &DBLibrary, nil
}
