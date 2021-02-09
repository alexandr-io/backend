package database

import (
	"context"
	"reflect"

	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ArraySubDocumentUpdate update a sub document in the database
//
// collection: The collection to work on
//
// filterDocument: The filter to find the first document
//
// filterSubDocument: The filter to find the sub document in the first document
//
// subDocumentPath: The path to the sub document
//
// update: The update to perform, if a field is not referenced, it will not be changed nor delete
//
// A data.NewHTTPErrorInfo is return if an error occurred
func ArraySubDocumentUpdate(ctx context.Context, collection *mongo.Collection, filterDocument bson.D, filterSubDocument bson.D, subDocumentPath string, update interface{}) error {
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{
		bson.D{{"$match", filterDocument}},
		bson.D{{"$unwind", bson.D{{"path", "$" + subDocumentPath}}}},
		bson.D{{"$match", filterSubDocument}},
		bson.D{{"$replaceRoot", bson.D{{"newRoot", "$" + subDocumentPath}}}},
		bson.D{{"$set", update}},
	})
	if err != nil {
		return err
	}

	reflectedData := reflect.ValueOf(update).Interface()

	if !cursor.Next(ctx) {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "The resource you requested does not exist.")
	}
	if err = cursor.Decode(&reflectedData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	finalFilter := filterDocument
	for _, e := range filterSubDocument {
		finalFilter = append(finalFilter, primitive.E{
			Key:   e.Key,
			Value: e.Value,
		})
	}

	_, err = collection.UpdateOne(ctx, finalFilter, bson.D{{"$set", bson.D{{subDocumentPath + ".$", reflectedData}}}})
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
