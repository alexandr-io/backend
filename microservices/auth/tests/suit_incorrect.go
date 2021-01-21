package tests

import "net/http"

const (
	duplicateSuit    = "Duplicate"
	noMatchSuit      = "No Match"
	invalidTokenSuit = "Invalid Token"
)

var wrongToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

var incorrectTests = []test{
	{
		TestSuit:         duplicateSuit,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/invitation/new" },
		AuthJWT:          nil,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    invitationEndFunction,
	},
	{
		TestSuit:   duplicateSuit,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/register" },
		AuthJWT:    nil,
		Body: userRegister{
			Email:           &randomEmail,
			Username:        &randomName,
			Password:        "test",
			ConfirmPassword: "test",
			InvitationToken: &invitationToken,
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:         duplicateSuit,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/invitation/" + invitationToken },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:   noMatchSuit,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/login" },
		AuthJWT:    nil,
		Body: userLogin{
			Login:    &randomName,
			Password: "wrong-password",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:   noMatchSuit,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/password/reset" },
		AuthJWT:    &authToken,
		Body: struct {
			Email string `json:"email"`
		}{
			Email: "wrong-email@test.test",
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:         invalidTokenSuit,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/auth" },
		AuthJWT:          &wrongToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:   invalidTokenSuit,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/refresh" },
		AuthJWT:    &wrongToken,
		Body: struct {
			RefreshToken string `json:"refresh_token"`
		}{
			RefreshToken: "wrong-token",
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecAuthIncorrectTests execute incorrect auth tests.
func ExecAuthIncorrectTests(environment string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}

	if err := execTestSuit(baseURL, incorrectTests); err != nil {
		return err
	}
	return nil
}
