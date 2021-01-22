package tests

import (
	"net/http"
)

const badRequestSuit = "Bad Request"

var badRequestTests = []test{
	{
		TestSuit:   badRequestSuit,
		HTTPMethod: http.MethodPut,
		URL:        func() string { return "/user" },
		AuthJWT:    &authToken,
		Body: struct {
			Username int `json:"username"`
		}{
			Username: 42,
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecUserBadRequestTests execute bad request user tests.
func ExecUserBadRequestTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	if err := execTestSuit(baseURL, badRequestTests); err != nil {
		return err
	}
	return nil
}
