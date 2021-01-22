package tests

import (
	"net/http"
)

const deleteSuit = "Delete"

var deleteTests = []test{
	{
		TestSuit:         deleteSuit,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/user" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:         deleteSuit,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/user" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecUserDeleteTests execute delete user tests.
func ExecUserDeleteTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	if err := execTestSuit(baseURL, deleteTests); err != nil {
		return err
	}
	return nil
}
