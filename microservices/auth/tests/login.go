package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// userLogin is the body parameter given to log in a user for test purpose
type userLogin struct {
	Login    *string `json:"login" validate:"required"`
	Password string  `json:"password" validate:"required"`
}

func loginEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var loginData user
	if err := json.Unmarshal(resBody, &loginData); err != nil {
		return err
	}
	// Store tokens for future uses
	authTokenLogin = loginData.AuthToken
	refreshTokenLogin = loginData.RefreshToken
	return nil
}

func loginLogoutSuiteEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var loginData user
	if err := json.Unmarshal(resBody, &loginData); err != nil {
		return err
	}
	// Store tokens for future uses
	authTokenLogout = loginData.AuthToken
	return nil
}
