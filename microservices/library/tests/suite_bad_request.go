package tests

import (
	"net/http"

	"github.com/alexandr-io/backend/library/data"
)

const badRequestSuite = "Bad Request"

var badRequestTests = []test{
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
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPut,
		URL:        func() string { return "/library" },
		AuthJWT:    &authToken,
		Body: struct {
			Name int `json:"name"`
		}{
			Name: 42,
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library" },
		AuthJWT:    &authToken,
		Body: data.Book{
			Title: libraryName,
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/book" },
		AuthJWT:    &authToken,
		Body: data.BookRetrieve{
			ID:        "42",
			LibraryID: "42",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPut,
		URL:        func() string { return "/book" },
		AuthJWT:    &authToken,
		Body: struct {
			ID        int    `json:"id"`
			LibraryID string `json:"libID"`
		}{
			ID:        42,
			LibraryID: "42",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/book" },
		AuthJWT:    &authToken,
		Body: bookCreation{
			Title: "Book name",
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
