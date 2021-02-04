package tests

import (
	"net/http"
)

const badRequestSuite = "Bad Request"

var badRequestTests = []test{
	{
		TestSuite:  badRequestSuite,
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

	if err := execTestSuite(baseURL, badRequestTests); err != nil {
		return err
	}
	return nil
}
