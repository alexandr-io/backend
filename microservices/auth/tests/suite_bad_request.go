package tests

import (
	"net/http"

	"github.com/alexandr-io/backend/auth/data"
)

const badRequestSuite = "Bad Request"

var badRequestTests = []test{
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/register" },
		AuthJWT:    nil,
		Body: userLogin{
			Login:    &randomName,
			Password: "test",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/login" },
		AuthJWT:    nil,
		Body: userRegister{
			Username: &randomName,
			Email:    &randomEmail,
			Password: "test",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/password/reset" },
		AuthJWT:    &authToken,
		Body: data.UserSendResetPasswordEmail{
			Email: "wrong-email",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecAuthBadRequestTests execute bad request auth tests.
func ExecAuthBadRequestTests(environment string, jwt string) error {
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
