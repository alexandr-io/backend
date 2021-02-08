package tests

import (
	"net/http"

	"github.com/alexandr-io/backend/library/data"
)

const badRequestSuite = "Bad Request"

var badRequestTests = []test{
	{
		TestSuite:        badRequestSuite,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/library" },
		AuthJWT:          &authToken,
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library" },
		AuthJWT:    &authToken,
		Body: data.Library{
			Description: libraryDescription,
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecLibraryBadRequestTests execute bad request library tests.
func ExecLibraryBadRequestTests(environment string, jwt string) error {
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
