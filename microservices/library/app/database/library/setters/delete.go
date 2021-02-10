package setters

import (
	"context"
	"errors"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"
	"github.com/alexandr-io/backend/library/database/libraries/getters"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/alexandr-io/backend/library/database/mongo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// DeleteLibrary delete the library for a user and the name of the library.
func DeleteLibrary(userID string, libraryID string) error {
	DBLibraries, err := getters.GetLibrariesFromUserID(userID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	found := false
	for _, DBLibrary := range DBLibraries.Libraries {
		if DBLibrary.ID == libraryID {
			found = true
			break
		}
	}
	if !found {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You do not have access to this library")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(library.Collection)

	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return err
	}
	libraryFilter := bson.D{{Key: "_id", Value: id}}
	deleteResult, err := collection.DeleteOne(ctx, libraryFilter)
	if err != nil {
		return err
	} else if deleteResult.DeletedCount == 0 {
		return errors.New("library does not exist")
	}

	for i, DBLibrary := range DBLibraries.Libraries {
		if DBLibrary.ID == libraryID {
			DBLibraries.Libraries = append(DBLibraries.Libraries[:i], DBLibraries.Libraries[i+1:]...)
			break
		}
	}

	collection = mongo.Instance.Db.Collection(libraries.Collection)

	userFilter := bson.D{{Key: "user_id", Value: userID}}

	updateValues := bson.D{{Key: "$set", Value: DBLibraries}}
	_ = collection.FindOneAndUpdate(ctx, userFilter, updateValues)
	return nil
}
