package tests

import (
	"github.com/alexandr-io/backend/library/data"
	"net/http"
)

const (
	workingSuit            = "Working"
	libraryName            = "Bookshelf"
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
		TestSuit:   workingSuit,
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
		TestSuit:         workingSuit,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/libraries" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    LibrariesGetEndFunction,
	},
	{
		TestSuit:   workingSuit,
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
		TestSuit:   workingSuit,
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
		TestSuit:   workingSuit,
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
		TestSuit:   workingSuit,
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
		TestSuit:   workingSuit,
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
		TestSuit:   workingSuit,
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

	if err := execTestSuit(baseURL, workingTests); err != nil {
		return err
	}
	return nil
}
