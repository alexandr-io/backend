package tests

import "net/http"

const logoutSuite = "Logout"

var authTokenLogout string

var logoutTests = []test{
	{
		TestSuite:  logoutSuite,
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
		CustomEndFunc: loginLogoutSuiteEndFunction,
	},
	{
		TestSuite:        logoutSuite,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/logout" },
		AuthJWT:          &authTokenLogout,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        logoutSuite,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/logout" },
		AuthJWT:          &authTokenLogout,
		Body:             nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        logoutSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/auth" },
		AuthJWT:          &authTokenLogout,
		Body:             nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecAuthLogoutTests execute logout auth tests.
func ExecAuthLogoutTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	if err := execTestSuite(baseURL, logoutTests); err != nil {
		return err
	}
	return nil
}
