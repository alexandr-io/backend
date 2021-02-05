package tests

import (
	"github.com/alexandr-io/backend/library/data"
	"net/http"
)

const (
	unauthorisedSuite = "Unauthorised"
	libraryInvalidID  = "invalid_id"
	bookInvalidID     = "invalid_id"
)

var unauthorisedTests = []test{
	{
		TestSuite:  unauthorisedSuite,
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
		CustomEndFunc: LibrayCreateEndFunction,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book" },
		AuthJWT:    &authToken,
		Body: bookCreation{
			Title:       bookTitle,
			Author:      bookAuthor,
			Publisher:   bookPublisher,
			Description: bookDescription,
		},
		ExpectedHTTPCode: http.StatusCreated,
		ExpectedResponse: data.Book{
			Title:       bookTitle,
			Author:      bookAuthor,
			Publisher:   bookPublisher,
			Description: bookDescription,
		},
		CustomEndFunc: BookCreateEndFunction,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodGet,
		URL:        func() string { return "/library/list" },
		AuthJWT:    nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library" },
		AuthJWT:    nil,
		Body: data.Library{
			Name:        libraryName,
			Description: libraryDescription,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc: nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodGet,
		URL:        func() string { return "/library/" + libraryInvalidID },
		AuthJWT:    &authToken,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodGet,
		URL:        func() string { return "/library/" + libraryID },
		AuthJWT:    nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library/" + libraryInvalidID },
		AuthJWT:    &authToken,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library/" + libraryID },
		AuthJWT:    nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryInvalidID + "/book" },
		AuthJWT:    &authToken,
		Body: bookCreation{
			Title:       bookTitle,
			Author:      bookAuthor,
			Publisher:   bookPublisher,
			Description: bookDescription,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book" },
		AuthJWT:    nil,
		Body: bookCreation{
			Title:       bookTitle,
			Author:      bookAuthor,
			Publisher:   bookPublisher,
			Description: bookDescription,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodGet,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID },
		AuthJWT:    nil,
		Body: bookRetrieve{
			ID:        &bookID,
			LibraryID: &libraryID,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodGet,
		URL:        func() string { return "/library/" + libraryInvalidID + "/book/" + bookID },
		AuthJWT:    &authToken,
		Body: bookRetrieve{
			ID:        &bookID,
			LibraryID: &libraryID,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodGet,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookInvalidID },
		AuthJWT:    &authToken,
		Body: bookRetrieve{
			ID:        &bookID,
			LibraryID: &libraryID,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID },
		AuthJWT:    nil,
		Body: data.Book{
			Description: bookDescriptionUpdated,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryInvalidID + "/book/" + bookID },
		AuthJWT:    &authToken,
		Body: data.Book{
			Description: bookDescriptionUpdated,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookInvalidID },
		AuthJWT:    &authToken,
		Body: data.Book{
			Description: bookDescriptionUpdated,
		},
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library/" + libraryID },
		AuthJWT:    nil,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library/" + libraryInvalidID },
		AuthJWT:    &authToken,
		ExpectedHTTPCode: http.StatusUnauthorized,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  unauthorisedSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library/" + libraryID },
		AuthJWT:    &authToken,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecLibraryUnauthorisedTests execute unauthorised library tests.
func ExecLibraryUnauthorisedTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	if err := execTestSuite(baseURL, unauthorisedTests); err != nil {
		return err
	}
	return nil
}
