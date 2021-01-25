package database

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"

	bson2 "github.com/globalsign/mgo/bson"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDataCreate creates an entry in mongodb for a user's book.
func UserDataCreate(ctx context.Context, userData data.UserData) (data.UserData, error) {
	collection := Instance.Db.Collection(CollectionBookUserData)

	// TODO: check if already exists
	_, err := collection.InsertOne(ctx, userData)
	if err != nil {
		return data.UserData{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	filter := bson.D{{"user_id", userData.UserID}}
	userDataRaw := collection.FindOne(ctx, filter)

	if err := userDataRaw.Decode(&userData); err != nil {
		return data.UserData{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return userData, nil
}

// ProgressRetrieve retrieves the user's progress from the mongo database
func ProgressRetrieve(ctx context.Context, progressRetrieve data.APIProgressRetrieve) (data.APIProgressData, error) {

	collection := Instance.Db.Collection(CollectionBookUserData)
	userFilter := bson.D{{"user_id", progressRetrieve.UserID}}
	userDataRaw := collection.FindOne(ctx, userFilter)
	var userData data.UserData

	if userDataRaw == nil {
		return data.APIProgressData{}, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Could not find user progress")
	}
	if err := userDataRaw.Decode(&userData); err != nil {
		return data.APIProgressData{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	for _, bookData := range userData.BookData {
		if bookData.BookID == progressRetrieve.BookID && bookData.LibraryID == progressRetrieve.LibraryID {
			return data.APIProgressData{
				BookID:       bookData.BookID,
				LibraryID:    bookData.LibraryID,
				Progress:     bookData.Progress,
				LastReadDate: bookData.LastReadDate,
			}, nil
		}
	}
	return data.APIProgressData{}, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Could not find progress for this book")
}

// ProgressUpdate updates the user's progress in the mongo database
func ProgressUpdate(ctx context.Context, progressData data.APIProgressData) (data.BookUserData, error) {

	collection := Instance.Db.Collection(CollectionBookUserData)
	userFilter := bson.D{{"user_id", progressData.UserID}}
	userDataRaw := collection.FindOne(ctx, userFilter)
	var userData data.UserData

	if userDataRaw == nil {
		return data.BookUserData{}, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Could not find user progress")
	}
	if err := userDataRaw.Decode(&userData); err != nil {
		return data.BookUserData{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	var returnValue data.BookUserData
	updated := false
	for i, bookData := range userData.BookData {
		if bookData.BookID == progressData.BookID && bookData.LibraryID == progressData.LibraryID {
			userData.BookData[i].Progress = progressData.Progress
			userData.BookData[i].LastReadDate = time.Now()

			returnValue = userData.BookData[i]
			updated = true
			break
		}
	}
	if !updated {
		newID, err := primitive.ObjectIDFromHex(bson2.NewObjectId().Hex())
		if err != nil {
			return returnValue, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}

		returnValue = data.BookUserData{
			ID:           newID,
			BookID:       progressData.BookID,
			LibraryID:    progressData.LibraryID,
			Progress:     progressData.Progress,
			LastReadDate: time.Now(),
		}
		userData.BookData = append(userData.BookData, returnValue)
	}

	_, err := collection.UpdateOne(ctx, userFilter, bson.D{{"$set", userData}})
	if err != nil {
		return returnValue, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return returnValue, nil
}
