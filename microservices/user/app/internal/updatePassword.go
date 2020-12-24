package internal

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdatePassword is the internal logic function used to update the password of an user.
func UpdatePassword(key string, message data.KafkaUpdatePassword) error {
	userObjectID, err := primitive.ObjectIDFromHex(message.ID)
	if err != nil {
		return producers.SendInternalErrorUserMessage(key, err.Error())
	}
	// Get the user from it's user ID
	user, err := database.GetUserByID(userObjectID)
	if err != nil {
		return producers.SendInternalErrorUpdatePasswordMessage(key, err.Error())
	}

	if _, err = database.UpdateUser(message.ID, data.User{
		Password: message.Password,
	}); err != nil {
		return producers.SendInternalErrorUpdatePasswordMessage(key, err.Error())
	}

	return producers.SendSuccessUpdatePasswordMessage(key, fiber.StatusOK, *user)
}
