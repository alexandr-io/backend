package user

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/gofiber/fiber/v2"
)

// FromLogin get an user by it's given login (username or email).
// In case of error, the proper error is set to the context and false is returned.
func FromLogin(login string) (*data.User, error) {
	if result, err := FromUsername(login); err == nil {
		return result, nil
	}
	if result, err := FromEmail(login); err == nil {
		return result, nil
	}

	return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "can't find user with login "+login)
}
