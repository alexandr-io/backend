package utils

import (
	"context"
	"errors"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"
	"github.com/alexandr-io/backend/library/database/mongo"
	"github.com/alexandr-io/berrors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// CheckLibrariesFieldDuplication check which field is a duplication.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
// The error returned is a formatted json of berrors.BadInput
func CheckLibrariesFieldDuplication(DBLibraries data.Libraries) error {
	errorsFields := make(map[string]string)

	filter := bson.D{{Key: "user_id", Value: DBLibraries.UserID}}
	object := &data.Libraries{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(libraries.Collection)
	if err := collection.FindOne(ctx, filter).Decode(object); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Check if the duplication is for the username field
	if object.UserID == DBLibraries.UserID {
		errorsFields["user_id"] = "This user already have a library."
	}
	if len(errorsFields) != 0 {
		return &data.BadInputError{
			JSONError: berrors.BadInputsJSON(errorsFields),
			Err:       errors.New("library duplication error"),
		}
	}
	return nil
}
