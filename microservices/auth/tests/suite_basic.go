package tests

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/alexandr-io/backend/tests/itgmtod"
)

const basicSuite = "Basic"

var (
	randomName      string
	randomEmail     string
	invitationToken string
	authToken       string
)

var basicTests = []test{
	{
		TestSuite:        basicSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/invitation/new" },
		AuthJWT:          nil,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    invitationEndFunction,
	},
	{
		TestSuite:  basicSuite,
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
		ExpectedHTTPCode: http.StatusCreated,
		ExpectedResponse: user{
			Username: &randomName,
			Email:    &randomEmail,
		},
		CustomEndFunc: registerEndFunction,
	},
	{
		TestSuite:        basicSuite,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/invitation/" + invitationToken },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        basicSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/auth" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: struct {
			Username *string `json:"username"`
		}{Username: &randomName},
		CustomEndFunc: nil,
	},
}

// ExecAuthBasicTests execute basic auth tests. Must be called first every time
func ExecAuthBasicTests(environment string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	baseURL, err := getBaseURL(environment)
	if err != nil {
		return "", err
	}
	randomName = itgmtod.RandStringRunes(12)
	randomEmail = randomName + "@test.test"
	if err := execTestSuite(baseURL, basicTests); err != nil {
		return "", err
	}
	return authToken, nil
}
