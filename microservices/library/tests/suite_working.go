package tests

import (
	"net/http"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
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
	groupID   string
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
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/library/" + libraryID },
		AuthJWT:          &authToken,
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
		URL:        func() string { return "/library/" + libraryID + "/book" },
		AuthJWT:    &authToken,
		Body: data.Book{
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
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodGet,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID },
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
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/library/" + libraryID + "/book/" + bookID + "/progress" },
		AuthJWT:          &authToken,
		Body:             data.APIProgressData{Progress: 42},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/library/" + libraryID + "/book/" + bookID + "/progress" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: data.APIProgressData{
			BookID:    bookID,
			LibraryID: libraryID,
			Progress:  42,
		},
		CustomEndFunc: nil,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodPost,
		URL:              func() string { return "/library/" + libraryID + "/book/" + bookID + "/progress" },
		AuthJWT:          &authToken,
		Body:             data.APIProgressData{Progress: 100},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/library/" + libraryID + "/book/" + bookID + "/progress" },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: data.APIProgressData{
			BookID:    bookID,
			LibraryID: libraryID,
			Progress:  100,
		},
		CustomEndFunc: nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodDelete,
		URL:        func() string { return "/library/" + libraryID + "/book/" + bookID },
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
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/group" },
		AuthJWT:    &authToken,
		Body: permissions.Group{
			Name:        "Test",
			Description: "Testing group, delete if in db",
			Priority:    0,
			Permissions: permissions.PermissionLibrary{
				Admin: typeconv.BoolPtr(true),
			},
		},
		ExpectedHTTPCode: http.StatusCreated,
		ExpectedResponse: nil,
		CustomEndFunc:    GroupPostEndFunction,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodGet,
		URL:              func() string { return "/library/" + libraryID + "/group/" + groupID },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: permissions.Group{
			Name:        "Test",
			Description: "Testing group, delete if in db",
			Priority:    0,
		},
		CustomEndFunc: nil,
	},
	{
		TestSuite:  workingSuite,
		HTTPMethod: http.MethodPost,
		URL:        func() string { return "/library/" + libraryID + "/group/" + groupID },
		AuthJWT:    &authToken,
		Body: permissions.Group{
			Name: "Test Updated",
		},
		ExpectedHTTPCode: http.StatusOK,
		ExpectedResponse: permissions.Group{
			Name:        "Test Updated",
			Description: "Testing group, delete if in db",
			Priority:    0,
		},
		CustomEndFunc: nil,
	},
	//{
	//	TestSuite: workingSuite,
	//	HTTPMethod: http.MethodGet,
	//	URL: func() string { return "/library/" + libraryID + "/user/groups"},
	//	AuthJWT: &authToken,
	//	Body: nil,
	//	ExpectedHTTPCode: http.StatusOK,
	//	ExpectedResponse: []permissions.Group{
	//		{
	//			Name:        "Test Updated",
	//			Description: "Testing group, delete if in db",
	//			Priority:    0,
	//		},
	//	},
	//	CustomEndFunc: nil,
	//},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/library/" + libraryID + "/group/" + groupID },
		AuthJWT:          &authToken,
		Body:             nil,
		ExpectedHTTPCode: http.StatusNoContent,
		ExpectedResponse: nil,
		CustomEndFunc:    nil,
	},
	{
		TestSuite:        workingSuite,
		HTTPMethod:       http.MethodDelete,
		URL:              func() string { return "/library/" + libraryID },
		AuthJWT:          &authToken,
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
