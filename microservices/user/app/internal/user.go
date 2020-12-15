package internal

import (
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the internal logic function used to get an user from an ID.
func User(key string, message string) error {
	// Create the bson objectID of the userID
	userObjectID, err := primitive.ObjectIDFromHex(message)
	if err != nil {
		return producers.SendInternalErrorUserMessage(key, err.Error())
	}

	// Get the user from it's user ID
	user, err := database.GetUserByID(userObjectID)
	if err != nil {
		if database.IsMongoNoDocument(err) {
			return producers.SendUnauthorizedUserMessage(key, err.Error())
		}
		return producers.SendInternalErrorUserMessage(key, err.Error())
	}

	return producers.SendSuccessUserMessage(key, fiber.StatusOK, *user)
}
