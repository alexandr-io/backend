package tests

import (
	"net/http"

	"github.com/alexandr-io/backend/tests/itgmtod"
)

const deleteSuite = "Delete"

var (
	randomName  string
	randomEmail string
)

var deleteTests = []test{
	{
		TestSuite:        deleteSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/user" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  deleteSuite,
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
	{
		TestSuite:        deleteSuite,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/user" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        deleteSuite,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/user" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecUserDeleteTests execute delete user tests.
func ExecUserDeleteTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	randomName = itgmtod.RandStringRunes(12)
	randomEmail = randomName + "@test.test"
	if err := execTestSuite(baseURL, deleteTests); err != nil {
		return err
	}
	return nil
}
