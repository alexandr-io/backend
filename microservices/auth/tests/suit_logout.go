package tests

import "net/http"

const logoutSuit = "Logout"

var authTokenLogout string

var logoutTests = []test{
	{
		TestSuit:   logoutSuit,
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
		CustomEndFunc: loginLogoutSuitEndFunction,
	},
	{
		TestSuit:         logoutSuit,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/logout" },
		AuthJWT:          &authTokenLogout,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:         logoutSuit,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/logout" },
		AuthJWT:          &authTokenLogout,
		Body:             nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:         logoutSuit,
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
func ExecAuthLogoutTests(environment string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}

	if err := execTestSuit(baseURL, logoutTests); err != nil {
		return err
	}
	return nil
}
