package internal

import (
	"context"
	"os"

	"github.com/alexandr-io/backend/common/generate"
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	grpcclient "github.com/alexandr-io/backend/user/grpc/client"
	"github.com/alexandr-io/backend/user/redis"
)

// Register is the internal logic function used to register a user.
func Register(ctx context.Context, newUser data.User) (*data.User, error) {
	// Insert the new data to the collection
	createdUser, err := user.Insert(data.User{
		Username:      newUser.Username,
		Email:         newUser.Email,
		Password:      newUser.Password,
		EmailVerified: false,
	})
	if err != nil {
		return nil, err
	}

	// Verify email
	verifyEmailToken := generate.RandomStringNoSpecialChar(12)
	if err := redis.StoreVerifyEmail(ctx, verifyEmailToken, newUser.Email); err != nil {
		return nil, err
	}
	verifyEmailURL := os.Getenv("USER_URI") + "/verify?token=" + verifyEmailToken
	if err := grpcclient.SendEmail(ctx, data.Email{
		Email:    newUser.Email,
		Username: newUser.Username,
		Type:     data.VerifyEmail,
		Data:     verifyEmailURL,
	}); err != nil {
		return nil, err
	}

	return createdUser, nil
}
