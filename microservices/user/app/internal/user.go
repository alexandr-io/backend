package internal

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"github.com/gofiber/fiber/v2"
)

// User is the internal logic function used to get an user from an ID.
func User(key string, message string) error {
	// Parse data
	kafkaUser, err := data.UnmarshalKafkaUser([]byte(message))
	if err != nil {
		return producers.SendInternalErrorUserMessage(key, err.Error())
	}

	var userDB *data.User

	if kafkaUser.ID != "" {
		// Get the user from it's user ID
		userDB, err = user.FromID(kafkaUser.ID)
		if err != nil {
			if database.IsMongoNoDocument(err) {
				return producers.SendUnauthorizedUserMessage(key, err.Error())
			}
			return producers.SendInternalErrorUserMessage(key, err.Error())
		}
	} else if kafkaUser.Email != "" {
		userDB, err = user.FromLogin(kafkaUser.Email)
		if err != nil {
			return producers.SendUnauthorizedUserMessage(key, err.Error())
		}
	} else {
		return producers.SendInternalErrorUserMessage(key, "No email nor user ID")
	}

	return producers.SendSuccessUserMessage(key, fiber.StatusOK, *userDB)
}
