package database

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/alexandr-io/backend/library/data"
)

// ProgressSpeedDB instance of ProgressSpeedCollection
var ProgressSpeedDB *ProgressSpeedCollection

// ProgressSpeedCollection contain the db collection for the progress speed interface
type ProgressSpeedCollection struct {
	collection *mongo.Collection
}

const progressSpeedCollectionName = "progress_speed"

// NewProgressSpeedCollection create a ProgressSpeedCollection
func NewProgressSpeedCollection(db *mongo.Database) *ProgressSpeedCollection {
	return &ProgressSpeedCollection{collection: db.Collection(progressSpeedCollectionName)}
}

func (c *ProgressSpeedCollection) Read(userID primitive.ObjectID, language string) (*data.ProgressSpeed, error) {
	filter := bson.D{
		{"user_id", userID},
		{"language", language},
	}

	var result data.ProgressSpeed
	if err := c.collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Progress speed not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &result, nil
}

// Upsert the user's progress speed in the database
// If no document is found, a new document is created
func (c *ProgressSpeedCollection) Upsert(progressSpeed *data.ProgressSpeed) error {
	filter := bson.D{
		{"user_id", progressSpeed.UserID},
		{"language", progressSpeed.Language},
	}

	if err := c.collection.FindOneAndUpdate(context.Background(), filter, bson.D{{"$set", progressSpeed}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1),
	).Decode(&progressSpeed); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
