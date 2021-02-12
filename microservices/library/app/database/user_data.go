package database

import (
	"context"

	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CollectionBookUserData is the name of the user data collection in mongodb
const CollectionBookUserData = "user_progress"

// ProgressRetrieve retrieves the user's book progress from the mongo database
func ProgressRetrieve(ctx context.Context, progressRetrieve data.APIProgressData) (*data.APIProgressData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)
	bookUserData, err := progressRetrieve.ToBookUserData()

	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	filter := bson.D{
		{"user_id", bookUserData.UserID},
		{"book_id", bookUserData.BookID},
		{"library_id", bookUserData.LibraryID},
	}
	var result data.BookUserData
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	progressData := result.ToAPIProgressData()
	return &progressData, nil
}

// ProgressUpdateOrInsert updates the user's book progress in the database
func ProgressUpdateOrInsert(ctx context.Context, bookUserData data.BookUserData) (*data.BookUserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)

	filter := bson.D{
		{"user_id", bookUserData.UserID},
		{"book_id", bookUserData.BookID},
		{"library_id", bookUserData.LibraryID},
	}

	if err := collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", bookUserData}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1),
	).Decode(&bookUserData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &bookUserData, nil
}

// progressDelete deletes the user's book progress data entry in the database
func progressDelete(ctx context.Context, bookUserData data.BookUserData) error {
	collection := Instance.Db.Collection(CollectionBookUserData)

	filter := bson.D{
		{"book_id", bookUserData.BookID},
		{"library_id", bookUserData.LibraryID},
	}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
