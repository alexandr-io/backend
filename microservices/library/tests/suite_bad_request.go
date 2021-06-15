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
			Name:        libraryName,
			Description: libraryDescription,
		},
		ExpectedHTTPCode: http.StatusCreated,
		ExpectedResponse: data.Library{
			Name:        libraryName,
			Description: libraryDescription,
		},
		CustomEndFunc: LibraryCreateEndFunction,
	},
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
	{
		TestSuite:        badRequestSuite,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/library/0/book/0/progress" },
		AuthJWT:          &authToken,
		Body:             data.BookProgressData{Progress: "progress"},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        badRequestSuite,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/library/" + libraryID + "/group" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        badRequestSuite,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/library/" + libraryID },
		AuthJWT:          &authToken,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID + "/data" },
		AuthJWT:    &authToken,
		Body: data.UserData{
			Type:    "highlight",
			Name:    "My data",
			Content: "Hello world",
			Tags:    nil,
			Offset:  "there",
			// OffsetEnd missing on purpose
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID + "/data" },
		AuthJWT:    &authToken,
		Body: data.UserData{
			Type:      "invalid",
			Name:      "My data",
			Content:   "Hello world",
			Tags:      nil,
			Offset:    "there",
			OffsetEnd: "",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID + "/data" },
		AuthJWT:    &authToken,
		Body: data.UserData{
			Type: "note",
			// Name missing on purpose
			Content: "Hello world",
			Offset:  "there",
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  badRequestSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID + "/data/" + dataID },
		AuthJWT:    &authToken,
		Body: data.UserData{
			Type:    "note",
			Name:    "Edited",
			Content: "yey",
			// missing Offset on purpose
		},
		ExpectedHTTPCode: http.StatusBadRequest,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        badRequestSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/library/" + libraryID + "/book/" + bookID + "/data/000000000000000000000000" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNotFound,
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
