package tests

import (
	"github.com/alexandr-io/backend/tests/integrationMethods"
	"math/rand"
	"net/http"
	"time"
)

const basicSuit = "Basic"

var (
	randomName      string
	randomEmail     string
	invitationToken string
	authToken       string
)

var basicTests = []test{
	{
		TestSuit:         basicSuit,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/invitation/new" },
		AuthJWT:          nil,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    invitationEndFunction,
	},
	{
		TestSuit:   basicSuit,
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
		TestSuit:         basicSuit,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/invitation/" + invitationToken },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:         basicSuit,
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
	randomName = integrationMethods.RandStringRunes(12)
	randomEmail = randomName + "@test.test"
	if err := execTestSuit(baseURL, basicTests); err != nil {
		return "", err
	}
	return authToken, nil
}
