package user

import (
	"errors"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/berrors"
)

// checkRegisterFieldDuplication check which field is a duplication on a register call.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
// The error returned is a formatted json of berrors.BadInput
func checkRegisterFieldDuplication(user data.User) error {
	errorsFields := make(map[string]string)

	if result, err := FromEmail(user.Email); err == nil && result.Email == user.Email {
		errorsFields["email"] = "Email has already been taken."
	}
	if result, err := FromUsername(user.Username); err == nil && result.Username == user.Username {
		errorsFields["username"] = "Username has already been taken."
	}

	if len(errorsFields) != 0 {
		return &data.BadInputError{
			JSONError: berrors.BadInputsJSON(errorsFields),
			Err:       errors.New("register duplication error"),
		}
	}
	return nil
}
