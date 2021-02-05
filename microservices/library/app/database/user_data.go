package database

import (
	"context"

	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CollectionBookUserData is the name of the user data collection in mongodb
const CollectionBookUserData = "book_user_data"

// userDataCreate creates an entry in mongodb for a user's book, if it doesn't exist.
func userDataCreate(ctx context.Context, userData data.UserData) (data.UserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)

	_, err := collection.InsertOne(ctx, userData)
	if err != nil {
		if !IsMongoDupKey(err) {
			return data.UserData{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
	}

	filter := bson.D{{"user_id", userData.UserID}}
	userDataRaw := collection.FindOne(ctx, filter)

	if err := userDataRaw.Decode(&userData); err != nil {
		return data.UserData{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return userData, nil
}

// ProgressRetrieve retrieves the user's progress from the mongo database
func ProgressRetrieve(ctx context.Context, progressRetrieve data.APIProgressRetrieve) (*data.APIProgressData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)
	userFilter := bson.D{{"user_id", progressRetrieve.UserID}}
	userDataRaw := collection.FindOne(ctx, userFilter)
	var userData data.UserData

	if err := userDataRaw.Decode(&userData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Could not find user progress")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	for _, bookData := range userData.BookData {
		if bookData.BookID == progressRetrieve.BookID && bookData.LibraryID == progressRetrieve.LibraryID {
			return &data.APIProgressData{
				BookID:       bookData.BookID,
				LibraryID:    bookData.LibraryID,
				Progress:     bookData.Progress,
				LastReadDate: bookData.LastReadDate,
			}, nil
		}
	}
	return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Could not find progress for this book")
}

// ProgressUpdate updates the user's progress in the mongo database
func ProgressUpdate(ctx context.Context, progressData data.APIProgressData) (*data.BookUserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)
	userFilter := bson.D{{"user_id", progressData.UserID}}
	bookFilter := bson.D{{"book_id", progressData.BookID}}

	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{
		bson.D{{"$match", userFilter}},
		bson.D{{"$unwind", bson.D{
			{"path", "$book_user_data"},
			{"preserveNullAndEmptyArrays", true}},
		}},
		bson.D{{"$replaceRoot", bson.D{
			{"newRoot", "$book_user_data"}},
		}},
		bson.D{{"$match", bookFilter}},
		bson.D{{"$set", bson.D{
			{"progress", progressData.Progress},
			{"last_read_date", progressData.LastReadDate},
		}},
		},
	})

	if err != nil {
		userDataRaw := collection.FindOne(ctx, userFilter)
		var userData data.UserData

		if err := userDataRaw.Decode(&userData); err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Could not find user progress")
			}
			return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}

		if len(userData.BookData) == 0 {
			return createBookData(ctx, progressData)
		}

		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if !cursor.Next(ctx) {
		return createBookData(ctx, progressData)
	}

	result := new(data.BookUserData)
	if err = cursor.Decode(result); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	filter := bson.D{
		{"user_id", progressData.UserID},
		{"book_user_data.book_id", progressData.BookID},
	}
	_, err = collection.UpdateOne(ctx, filter, bson.D{
		{"$set", bson.D{{"book_user_data.$", result}}},
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return result, nil
}

func createBookData(ctx context.Context, progressData data.APIProgressData) (*data.BookUserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)
	userFilter := bson.D{{"user_id", progressData.UserID}}
	result := progressData.ToBookUserData()
	_, err := collection.UpdateOne(ctx, userFilter, bson.D{{"$push", bson.D{{"book_user_data", result}}}})

	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &result, nil
}
