package tests

import (
	"github.com/alexandr-io/backend/tests/itgmtod"
	"net/http"
)

const workingSuit = "Working"

var (
	randomName  string
	randomEmail string
)

var workingTests = []test{
	{
		TestSuit:         workingSuit,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/user" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuit:   workingSuit,
		HTTPMethod: http.MethodPut,
		URL:        func() string { return "/user" },
		AuthJWT:    &authToken,
		Body: user{
			Username: &randomName,
			Email:    &randomEmail,
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: user{
			Username: &randomName,
			Email:    &randomEmail,
		},
		CustomEndFunc: nil,
	},
}

// ExecUserWorkingTests execute working user tests.
func ExecUserWorkingTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	randomName = itgmtod.RandStringRunes(12)
	randomEmail = randomName + "@test.test"
	if err := execTestSuit(baseURL, workingTests); err != nil {
		return err
	}
	return nil
}
