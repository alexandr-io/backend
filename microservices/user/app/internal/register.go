package internal

import (
	"errors"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"github.com/gofiber/fiber/v2"
)

// Register is the internal logic function used to register a user.
func Register(key string, message data.KafkaUserRegisterRequest) error {
	// Insert the new data to the collection

	createdUser, err := user.Insert(data.User{
		Username: message.Username,
		Email:    message.Email,
		Password: message.Password,
	})
	if err != nil {
		var badInput *data.BadInputError
		if errors.As(err, &badInput) {
			return producers.SendBadRequestRegisterMessage(key, badInput.JSONError)
		}
		return producers.SendInternalErrorRegisterMessage(key, err.Error())
	}
	return producers.SendSuccessRegisterMessage(key, fiber.StatusCreated, *createdUser)
}
