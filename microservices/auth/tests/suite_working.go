package tests

import (
	"net/http"

	"github.com/alexandr-io/backend/auth/data"
)

const workingSuite = "Working"

var (
	authTokenLogin    string
	refreshTokenLogin string
	authTokenRefresh  string
)

var workingTests = []test{
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/login" },
		AuthJWT:    nil,
		Body: userLogin{
			Login:    &randomName,
			Password: "test",
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: user{
			Username: &randomName,
			Email:    &randomEmail,
		},
		CustomEndFunc: loginEndFunction,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/auth" },
		AuthJWT:          &authTokenLogin,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: struct {
			Username *string `json:"username"`
		}{Username: &randomName},
		CustomEndFunc: nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/refresh" },
		AuthJWT:    &authTokenLogin,
		Body: struct {
			RefreshToken *string `json:"refresh_token"`
		}{
			RefreshToken: &refreshTokenLogin,
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: user{
			Username: &randomName,
			Email:    &randomEmail,
		},
		CustomEndFunc: refreshEndFunction,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/auth" },
		AuthJWT:          &authTokenRefresh,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: struct {
			Username *string `json:"username"`
		}{
			Username: &randomName,
		},
		CustomEndFunc: nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/password/reset" },
		AuthJWT:    &authToken,
		Body: struct {
			Email *string `json:"email"`
		}{
			Email: &randomEmail,
		},
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPut,
		URL:        func() string { return "/password/update" },
		AuthJWT:    &authToken,
		Body: data.UpdatePassword{
			CurrentPassword: "test",
			NewPassword:     "newPassword",
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: data.User{
			Username: randomName,
			Email:    randomEmail,
		},
		CustomEndFunc: nil,
	},
}

// ExecAuthWorkingTests execute working auth tests.
func ExecAuthWorkingTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	if err := execTestSuite(baseURL, workingTests); err != nil {
		return err
	}
	return nil
}
