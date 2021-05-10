package user

import (
	"context"
	"errors"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/berrors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert insert a new user into the database.
func Insert(user data.User) (*data.User, error) {

	insertedResult, err := database.UserCollection.InsertOne(context.Background(), user)
	if database.IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		e := checkRegisterFieldDuplication(user)
		var badInput *data.BadInputError
		if errors.As(e, &badInput) {
			return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, badInput.Error())
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	user.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &user, nil
}

// checkRegisterFieldDuplication check which field is a duplication on a register call.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
// The error returned is a formatted json of berrors.BadInput
func checkRegisterFieldDuplication(user data.User) error {
	errorsFields := make(map[string]string)

	if result, err := FromEmail(user.Email); err == nil && result.Email == user.Email {
		errorsFields["email"] = "Email has already been taken."
	}
	if result, err := FromUsername(user.Username); err == nil && result.Username == user.Username {
		errorsFields["username"] = "Username has already been taken."
	}

	if len(errorsFields) != 0 {
		return &data.BadInputError{
			JSONError: berrors.BadInputsJSON(errorsFields),
			Err:       errors.New("register duplication error"),
		}
	}
	return nil
}
