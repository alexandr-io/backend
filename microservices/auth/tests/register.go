package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// userRegister is the Body parameter given to register a new userOld to the database.
// test copy with email and username set to pointers.
type userRegister struct {
	Email           *string `json:"email" validate:"required,email"`
	Username        *string `json:"username" validate:"required"`
	Password        string  `json:"password" validate:"required"`
	ConfirmPassword string  `json:"confirm_password" validate:"required"`
	InvitationToken *string `json:"invitation_token,omitempty" validate:"required,len=10"`
}

func registerEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var registerData user
	if err := json.Unmarshal(resBody, &registerData); err != nil {
		return err
	}
	// Store tokens for future uses
	authToken = registerData.AuthToken
	return nil
}
