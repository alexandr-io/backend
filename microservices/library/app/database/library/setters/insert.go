package setters

import (
	"context"
	"errors"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database/libraries"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/alexandr-io/backend/library/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// InsertLibrary insert on the database a new library in the user's libraries.
func InsertLibrary(userID string, DBLibrary data.Library) (*data.Libraries, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userFilter := bson.D{{Key: "user_id", Value: userID}}

	collection := mongo.Instance.Db.Collection(library.Collection)

	insertedResult, err := collection.InsertOne(ctx, DBLibrary)
	if err != nil {
		return nil, err
	}

	collection = mongo.Instance.Db.Collection(libraries.Collection)

	object := &data.Libraries{}
	if err := collection.FindOne(ctx, userFilter).Decode(object); err != nil {
		return object, err
	}

	id, ok := insertedResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("can't cast InsertedID to primitive.ObjectID")
	}
	object.Libraries = append(object.Libraries, data.LibraryData{
		ID:          id.Hex(),
		Permissions: []permissions.Permission{"owner"},
	})

	updateValues := bson.D{{Key: "$set", Value: object}}
	findAndUpdateOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)
	updatedSingleResult := collection.FindOneAndUpdate(ctx, userFilter, updateValues, findAndUpdateOptions)
	if err := updatedSingleResult.Decode(object); err != nil {
		return object, err
	}
	return object, nil
}
