package jwt

import (
	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/kafka"

	"github.com/gofiber/fiber/v2"
)

// GetUserFromContextJWT return a user from a jwt contained in the fiber context
func GetUserFromContextJWT(ctx *fiber.Ctx) (*data.User, error) {
	// extract user ID from JWT
	userID, err := ExtractJWTUserID(ctx)
	if err != nil {
		return nil, err
	}

	// Get the user from user MS using kafka
	kafkaUser, err := kafka.UserRequestHandler(userID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &data.User{
		ID:       kafkaUser.ID,
		Username: kafkaUser.Username,
		Email:    kafkaUser.Email,
	}, nil
}