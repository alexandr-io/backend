package internal

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
)

// Register is the internal logic function used to register a user.
func Register(newUser data.User) (*data.User, error) {
	// Insert the new data to the collection
	createdUser, err := user.Insert(data.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
