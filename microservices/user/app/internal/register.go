package internal

import (
	"errors"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"github.com/gofiber/fiber/v2"
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
	createdUser, err := database.GetUserByID(insertedResult.InsertedID)
	if err != nil {
		return producers.SendInternalErrorRegisterMessage(key, err.Error())
	}

	return producers.SendSuccessRegisterMessage(key, fiber.StatusCreated, *createdUser)
}
