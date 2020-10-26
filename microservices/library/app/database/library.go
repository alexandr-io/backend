package database

import (
	"context"
	"errors"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/berrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// CollectionLibrary is the name of the library collection in mongodb
const CollectionLibrary = "library"

//
// Getters
//

// GetLibraryByUsername get a library by the username.
// In case of error, the proper error is set to the context and false is returned.
func GetLibraryByUsername(username string) (*data.Library, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibrary)
	object := &data.Library{}

	// Get library by username
	usernameFilter := bson.D{{Key: "username", Value: username}}
	filteredByUsernameSingleResult := collection.FindOne(ctx, usernameFilter)
	// Return the library object if library is found
	if err := filteredByUsernameSingleResult.Decode(object); err == nil {
		return object, err
	}

	// Return the library object
	return object, nil
}

//
// Setters
//

// InsertLibrary insert a new library into the database.
func InsertLibrary(library data.Library) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibrary)
	insertedResult, err := collection.InsertOne(ctx, data.Library{
		Username: library.Username,
	})
	if IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		err := checkLibraryFieldDuplication(library)
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return insertedResult, nil
}

// checkLibraryFieldDuplication check which field is a duplication.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
// The error returned is a formatted json of berrors.BadInput
func checkLibraryFieldDuplication(library data.Library) error {
	errorsFields := make(map[string]string)

	// Check if the duplication is for the username field
	foundLibrary, err := GetLibraryByUsername(library.Username)
	if err == nil && foundLibrary.Username == library.Username {
		errorsFields["username"] = "You already have a library."
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
