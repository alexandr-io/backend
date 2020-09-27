package internal

import (
	"errors"
	"net/http"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/kafka/producers"
)

// Register is the internal logic function used to register a user.
func Register(message data.KafkaUserRegisterRequest) error {
	// Insert the new data to the collection
	insertedResult, err := database.InsertUserRegister(data.User{
		Username: message.Data.Username,
		Email:    message.Data.Email,
		Password: message.Data.Password,
	})
	if err != nil {
		var badInput *data.BadInput
		if errors.As(err, &badInput) {
			return producers.SendBadRequestRegisterMessage(message.UUID, badInput.JSONError)
		} else {
			return producers.SendInternalErrorRegisterMessage(message.UUID, err.Error())
		}
	}

	// Get the newly created user
	createdUser, ok := database.GetUserByID(insertedResult.InsertedID)
	if !ok {
		return producers.SendInternalErrorRegisterMessage(message.UUID, "internal server error after user insertion")
	}

	return producers.SendSuccessRegisterMessage(message.UUID, http.StatusCreated, *createdUser)
}
