package database

import (
	"context"

	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CollectionBookUserData is the name of the user data collection in mongodb
const CollectionBookUserData = "user_progress"

// ProgressRetrieve retrieves the user's progress from the mongo database
func ProgressRetrieve(ctx context.Context, progressRetrieve data.APIProgressRetrieve) (*data.BookUserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)
	filter := bson.D{
		{"user_id", progressRetrieve.UserID},
		{"book_id", progressRetrieve.BookID},
		{"library_id", progressRetrieve.LibraryID},
	}

	result := new(data.BookUserData)
	err := collection.FindOne(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User data not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return result, nil
}

// ProgressUpdate updates the user's progress in the mongo database
func ProgressUpdate(ctx context.Context, apiProgressData data.APIProgressData) (*data.BookUserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)
	bookUserData, err := apiProgressData.ToBookUserData()

	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	filter := bson.D{
		{"user_id", bookUserData.UserID},
		{"book_id", bookUserData.BookID},
		{"library_id", bookUserData.LibraryID},
	}

	_, err = collection.UpdateOne(ctx, filter, bson.D{{"$set", bookUserData}})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return bookUserData, nil
}

func progressCreate(ctx context.Context, progressData data.APIProgressData) (*data.BookUserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)
	bookUserData, err := progressData.ToBookUserData()

	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	_, err = collection.InsertOne(ctx, bookUserData)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return bookUserData, nil
}

func progressDelete(ctx context.Context, progressData data.APIProgressData) error {
	collection := Instance.Db.Collection(CollectionBookUserData)

	bookUserData, err := progressData.ToBookUserData()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	filter := bson.D{
		{"user_id", bookUserData.UserID},
		{"book_id", bookUserData.BookID},
		{"library_id", bookUserData.LibraryID},
	}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
