package internal

import (
	"errors"
	"net/http"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/kafka/producers"
)

// Register is the internal logic function used to register a user.
func Register(key string, message data.KafkaUserRegisterRequest) error {
	// Insert the new data to the collection
	insertedResult, err := database.InsertUserRegister(data.User{
		Username: message.Data.Username,
		Email:    message.Data.Email,
		Password: message.Data.Password,
	})
	if err != nil {
		var badInput *data.BadInput
		if errors.As(err, &badInput) {
			return producers.SendBadRequestRegisterMessage(key, badInput.JSONError)
		}
		return producers.SendInternalErrorRegisterMessage(key, err.Error())
	}

	// Get the newly created user
	createdUser, ok := database.GetUserByID(insertedResult.InsertedID)
	if !ok {
		return producers.SendInternalErrorRegisterMessage(key, "internal server error after user insertion")
	}

	return producers.SendSuccessRegisterMessage(key, http.StatusCreated, *createdUser)
}
