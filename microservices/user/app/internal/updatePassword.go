package internal

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"github.com/gofiber/fiber/v2"
)

// UpdatePassword is the internal logic function used to update the password of an user.
func UpdatePassword(key string, message data.KafkaUpdatePassword) error {
	userDB, err := user.Update(message.ID, data.User{
		Password: message.Password,
	})
	if err != nil {
		return producers.SendInternalErrorUpdatePasswordMessage(key, err.Error())
	}
	return producers.SendSuccessUpdatePasswordMessage(key, fiber.StatusOK, *userDB)
}
