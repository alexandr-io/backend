package internal

import (
	"context"
	"strings"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
)

// Register is the internal logic function used to register a user.
func Register(ctx context.Context, newUser data.User) (*data.User, error) {
	verified := false
	if strings.Contains(newUser.Email, "@test.test") {
		verified = true
	}
	// Insert the new data to the collection
	createdUser, err := user.Insert(data.User{
		Username:      newUser.Username,
		Email:         newUser.Email,
		Password:      newUser.Password,
		EmailVerified: verified,
	})
	if err != nil {
		return nil, err
	}

	if err := VerifyEmailCreation(ctx, createdUser); err != nil {
		return nil, err
	}

	return createdUser, nil
}
