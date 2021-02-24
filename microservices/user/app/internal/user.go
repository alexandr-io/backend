package internal

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/gofiber/fiber/v2"
)

// User is the internal logic function used to get an user from an ID.
func User(ID string, email string) (*data.User, error) {

	var userData *data.User
	var err error = nil
	if ID != "" {
		userData, err = user.FromID(ID)
	} else if email != "" {
		userData, err = user.FromEmail(email)
	} else {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, "no ID nor email received")
	}
	if err != nil {
		return nil, err
	}

	return userData, nil
}
