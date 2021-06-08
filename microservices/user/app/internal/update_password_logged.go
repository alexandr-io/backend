package internal

import (
	"context"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	grpcclient "github.com/alexandr-io/backend/user/grpc/client"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdatePasswordLogged is the internal logic function used to update the password of a logged user.
func UpdatePasswordLogged(ctx context.Context, id primitive.ObjectID, currentPassword string, newPassword string) (*data.User, error) {
	userData, err := user.FromID(id)
	if err != nil {
		return nil, err
	}

	if !comparePasswords(userData.Password, []byte(currentPassword)) {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Unauthorized")
	}
	if !userData.EmailVerified {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Email must be verified to change your password")
	}

	userData, err = user.Update(id, data.User{
		Password: newPassword,
	})
	if err != nil {
		return nil, err
	}

	grpcclient.SendEmail(ctx, data.Email{
		Email:    userData.Email,
		Username: userData.Username,
		Type:     data.UpdatedPassword,
	})
	return userData, nil
}
