package internal

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
)

// UpdatePassword is the internal logic function used to update the password of an user.
func UpdatePassword(id string, password string) (*data.User, error) {
	userData, err := user.Update(id, data.User{
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return userData, nil
}
