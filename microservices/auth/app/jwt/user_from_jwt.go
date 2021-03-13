package jwt

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"
	grpcclient "github.com/alexandr-io/backend/auth/grpc/client"

	"github.com/gofiber/fiber/v2"
)

// GetUserFromContextJWT return a user from a jwt contained in the fiber context
func GetUserFromContextJWT(ctx *fiber.Ctx) (*data.User, error) {
	// extract user ID from JWT
	userID, err := ExtractJWTUserID(ctx)
	if err != nil {
		return nil, err
	}

	// Get the user from user MS using grpc
	userData, err := grpcclient.User(context.Background(), data.User{ID: userID})
	if err != nil {
		return nil, err
	}
	return userData, nil
}
