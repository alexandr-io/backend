package tests

import (
	"net/http"

	"github.com/alexandr-io/backend/library/data"
)

const (
	workingSuite           = "Working"
	libraryName            = "Library-test"
	libraryDescription     = "My bookshelf"
	bookTitle              = "Memoirs of Napoleon Bonaparte"
	bookAuthor             = "Louis Antoine Fauvelet de Bourrienne"
	bookPublisher          = "Public domain in the USA"
	bookDescription        = "Translated from: Mémoires sur Napoléon, le Directoire le Consulat, l'Empire et la Restauration."
	bookDescriptionUpdated = "Biographie de Napoleon Bonaparte"
)

var (
	authToken string
	libraryID string
	bookID    string
)

var workingTests = []test{
	{
		TestSuite:  workingSuite,
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
		CustomEndFunc: nil,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/libraries" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    LibrariesGetEndFunction,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPut,
		URL:        func() string { return "/library" },
		AuthJWT:    &authToken,
		Body: data.LibraryName{
			Name: libraryName,
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: data.Library{
			Name:        libraryName,
			Description: libraryDescription,
		},
		CustomEndFunc: nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/book" },
		AuthJWT:    &authToken,
		Body: bookCreation{
			Title:       bookTitle,
			Author:      bookAuthor,
			Publisher:   bookPublisher,
			Description: bookDescription,
			LibraryID:   &libraryID,
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
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPut,
		URL:        func() string { return "/book" },
		AuthJWT:    &authToken,
		Body: bookRetrieve{
			ID:        &bookID,
			LibraryID: &libraryID,
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: data.Book{
			Title:       bookTitle,
			Author:      bookAuthor,
			Publisher:   bookPublisher,
			Description: bookDescription,
		},
		CustomEndFunc: nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID },
		AuthJWT:    &authToken,
		Body: data.Book{
			Description: bookDescriptionUpdated,
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: data.Book{
			Title:       bookTitle,
			Author:      bookAuthor,
			Publisher:   bookPublisher,
			Description: bookDescriptionUpdated,
		},
		CustomEndFunc: nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/book" },
		AuthJWT:    &authToken,
		Body: bookRetrieve{
			ID:        &bookID,
			LibraryID: &libraryID,
		},
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library" },
		AuthJWT:    &authToken,
		Body: data.LibraryName{
			Name: libraryName,
		},
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
}

// ExecLibraryWorkingTests execute working library tests.
func ExecLibraryWorkingTests(environment string, jwt string) error {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}
	authToken = jwt

	if err := execTestSuite(baseURL, workingTests); err != nil {
		return err
	}
	return nil
}