package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/berrors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CollectionLibraries is the name of the libraries collection in mongodb
const CollectionLibraries = "libraries"

// CollectionLibrary is the name of the library collection in mongodb
const CollectionLibrary = "library"

//
// Getters
//

// FindOneWithFilter fill the given object with a mongodb single result filtered by the given filters.
func FindOneWithFilter(collectionName string, object interface{}, filters interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(collectionName)
	filteredSingleResult := collection.FindOne(ctx, filters)
	return filteredSingleResult.Decode(object)
}

// GetLibraryByID retrieve a library from its ID
func GetLibraryByID(libraryID string) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibrary)

	library := &data.Library{}

	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	libraryFilter := bson.D{{Key: "_id", Value: id}}

	filteredByLibraryIDSingleResult := collection.FindOne(ctx, libraryFilter)
	if err := filteredByLibraryIDSingleResult.Decode(library); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return library, err
}

// GetLibraryByUserIDAndLibraryID retrieve a library from a data.LibraryOwner and a library ID if the user have access to the library
func GetLibraryByUserIDAndLibraryID(user data.LibrariesOwner, libraryID string) (*data.Library, error) {
	libraries, err := GetLibrariesNamesByUserID(user)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	found := false
	for _, library := range libraries.Libraries {
		if library.ID == libraryID {
			found = true
			break
		}
	}
	if !found {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You do not have access to this library")
	}

	return GetLibraryByID(libraryID)
}

// GetLibraryByUserIDAndName retrieve a library from a user and the name of a library.
func GetLibraryByUserIDAndName(user data.LibrariesOwner, library data.LibraryName) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibraries)

	object := &data.Library{}

	usernameFilter := bson.D{{Key: "user_id", Value: user.UserID}}
	projection := bson.D{{"libraries", true}}

	filterOptions := options.FindOne().SetProjection(projection)
	filteredByUsernameSingleResult := collection.FindOne(ctx, usernameFilter, filterOptions)

	libraries := &data.Libraries{}
	// Return the library object if library is found
	if err := filteredByUsernameSingleResult.Decode(libraries); err != nil {
		return object, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	collection = Instance.Db.Collection(CollectionLibrary)
	for _, libraryID := range libraries.Libraries {

		currentLibrary := &data.Library{}

		id, err := primitive.ObjectIDFromHex(libraryID)
		if err != nil {
			return object, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
		libraryFilter := bson.D{{Key: "_id", Value: id}}

		filteredByLibraryIDSingleResult := collection.FindOne(ctx, libraryFilter)
		if err := filteredByLibraryIDSingleResult.Decode(currentLibrary); err != nil {
			return object, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
		if currentLibrary.Name == library.Name {
			return currentLibrary, nil
		}
	}

	return object, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Library does not exist")
}

// GetLibrariesByUsername get the libraries the current user has access to.
// In case of error, the proper error is set to the context and false is returned.
func GetLibrariesByUsername(user data.LibrariesOwner) (*data.Libraries, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibraries)
	object := &data.Libraries{}

	usernameFilter := bson.D{{Key: "user_id", Value: user.UserID}}
	filteredByUsernameSingleResult := collection.FindOne(ctx, usernameFilter)
	if err := filteredByUsernameSingleResult.Decode(object); err == nil {
		return object, err
	}

	// Return the libraries object
	return object, nil
}

// GetLibrariesNamesByUserID get a list of the user libraries name.
// In case of error, the proper error is set to the context and false is returned.
func GetLibrariesNamesByUserID(user data.LibrariesOwner) (*data.LibrariesNames, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibraries)

	// Get library by username
	usernameFilter := bson.D{{Key: "user_id", Value: user.UserID}}
	projection := bson.D{{"libraries", true}}

	filterOptions := options.FindOne().SetProjection(projection)
	filteredByUsernameSingleResult := collection.FindOne(ctx, usernameFilter, filterOptions)

	libraries := &data.Libraries{}
	librariesNames := &data.LibrariesNames{}
	// Return the library object if library is found
	if err := filteredByUsernameSingleResult.Decode(libraries); err != nil {
		return librariesNames, err
	}

	collection = Instance.Db.Collection(CollectionLibrary)
	for _, libraryID := range libraries.Libraries {

		library := &data.Library{}

		id, err := primitive.ObjectIDFromHex(libraryID)
		if err != nil {
			return librariesNames, err
		}
		libraryFilter := bson.D{{Key: "_id", Value: id}}
		projection := bson.D{{"name", true}}

		filterOptions := options.FindOne().SetProjection(projection)
		filteredByLibraryIDSingleResult := collection.FindOne(ctx, libraryFilter, filterOptions)
		if err := filteredByLibraryIDSingleResult.Decode(library); err != nil {
			return librariesNames, err
		}
		librariesNames.Libraries = append(librariesNames.Libraries, data.LibraryName{
			ID:   libraryID,
			Name: library.Name,
		})
	}

	// Return the library object
	return librariesNames, nil
}

//
// Setters
//

// InsertLibraries create the libraries of an user.
func InsertLibraries(libraries data.Libraries) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibraries)

	insertedResult, err := collection.InsertOne(ctx, libraries)
	if IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		err := checkLibrariesFieldDuplication(libraries)
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return insertedResult, nil
}

// InsertLibrary insert on the database a new library in the user's libraries.
func InsertLibrary(user data.LibrariesOwner, library data.Library) (*data.Libraries, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userFilter := bson.D{{Key: "user_id", Value: user.UserID}}

	collection := Instance.Db.Collection(CollectionLibrary)

	err := checkLibraryFieldDuplication(user, library)
	if err != nil {
		return nil, err
	}

	insertedResult, err := collection.InsertOne(ctx, library)
	if err != nil {
		return nil, err
	}

	collection = Instance.Db.Collection(CollectionLibraries)
	object := &data.Libraries{}

	filteredByUsernameSingleResult := collection.FindOne(ctx, userFilter)

	if err := filteredByUsernameSingleResult.Decode(object); err != nil {
		return object, err
	}

	id, ok := insertedResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("can't cast InsertedID to primitive.ObjectID")
	}
	object.Libraries = append(object.Libraries, id.Hex())

	updateValues := bson.D{{Key: "$set", Value: object}}
	findAndUpdateOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)
	updatedSingleResult := collection.FindOneAndUpdate(ctx, userFilter, updateValues, findAndUpdateOptions)
	if err := updatedSingleResult.Decode(object); err != nil {
		return object, err
	}
	return object, nil
}

// DeleteLibrary delete the library for a user and the name of the library.
func DeleteLibrary(user data.LibrariesOwner, libraryID string) error {
	libraries, err := GetLibrariesByUsername(data.LibrariesOwner{UserID: user.UserID})
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	found := false
	for _, library := range libraries.Libraries {
		if library == libraryID {
			found = true
			break
		}
	}
	if !found {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You do not have access to this library")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibrary)

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

	for i, libraryElemID := range libraries.Libraries {
		if libraryElemID == libraryID {
			libraries.Libraries = append(libraries.Libraries[:i], libraries.Libraries[i+1:]...)
			break
		}
	}

	collection = Instance.Db.Collection(CollectionLibraries)

	userFilter := bson.D{{Key: "user_id", Value: user.UserID}}

	updateValues := bson.D{{Key: "$set", Value: libraries}}
	_ = collection.FindOneAndUpdate(ctx, userFilter, updateValues)
	return nil
}

// checkLibrariesFieldDuplication check which field is a duplication.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
// The error returned is a formatted json of berrors.BadInput
func checkLibrariesFieldDuplication(libraries data.Libraries) error {
	errorsFields := make(map[string]string)

	filter := bson.D{{Key: "user_id", Value: libraries.UserID}}
	object := &data.Libraries{}

	err := FindOneWithFilter(CollectionLibrary, object, filter)
	// Check if the duplication is for the username field
	if err == nil && object.UserID == libraries.UserID {
		errorsFields["user_id"] = "This user already have a library."
	} else if err != nil {
		log.Println(err)
		return err
	}

	if len(errorsFields) != 0 {
		return &data.BadInputError{
			JSONError: berrors.BadInputsJSON(errorsFields),
			Err:       errors.New("library duplication error"),
		}
	}
	return nil
}

// checkLibraryFieldDuplication check which field is a duplication.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
// The error returned is a formatted json of berrors.BadInput
func checkLibraryFieldDuplication(user data.LibrariesOwner, library data.Library) error {
	errorsFields := make(map[string]string)

	// Check if the duplication is for the username field
	foundLibraries, err := GetLibrariesNamesByUserID(user)
	if err == nil {
		for _, currentLibrary := range foundLibraries.Libraries {
			if currentLibrary.Name == library.Name {
				errorsFields["name"] = "You already have a library with this name."
				break
			}
		}
	} else {
		return err
	}

	if len(errorsFields) != 0 {
		return &data.BadInputError{
			JSONError: berrors.BadInputsJSON(errorsFields),
			Err:       errors.New("library duplication error"),
		}
	}
	return nil
}
