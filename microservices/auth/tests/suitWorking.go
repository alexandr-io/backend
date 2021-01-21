package tests

import (
	"net/http"
)

const workingSuit = "Working"

var (
	authTokenLogin    string
	refreshTokenLogin string
	authTokenRefresh  string
)

var workingTests = []test{
	{
		TestSuit:   workingSuit,
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
		TestSuit:         workingSuit,
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
		TestSuit:   workingSuit,
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
		TestSuit:         workingSuit,
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
		TestSuit:   workingSuit,
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
}

// ExecAuthWorkingTests execute working auth tests.
func ExecAuthWorkingTests(environment string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}

	if err := execTestSuit(baseURL, workingTests); err != nil {
		return err
	}
	return nil
}
