package tests

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/alexandr-io/backend/tests/itgmtod"
	"net/http"
)

var authToken string

type test struct {
	TestSuit         string
	HTTPMethod       string
	URL              func() string
	AuthJWT          *string
	Body             interface{}
	ExpectedHTTPCode int
	ExpectedResponse interface{}
	CustomEndFunc    func(*http.Response) error
}

func execTestSuit(baseURL string, testSuite []test) error {
	for _, currentTest := range testSuite {
		url := itgmtod.JoinURL(baseURL, currentTest.URL())
		body, err := itgmtod.MarshallBody(currentTest.Body)
		if err != nil {
			newFailureMessage(currentTest.HTTPMethod, url, currentTest.TestSuit, err.Error())
			return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
		}

		// Test the route
		res, err := itgmtod.TestRoute(
			currentTest.HTTPMethod,
			url,
			currentTest.AuthJWT,
			bytes.NewReader(body),
			currentTest.ExpectedHTTPCode)
		if err != nil {
			newFailureMessage(currentTest.HTTPMethod, currentTest.URL(), currentTest.TestSuit, err.Error())
			return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
		}
		// Check expected response Body
		if currentTest.ExpectedResponse != nil {
			if err := itgmtod.CheckExpectedResponse(res, currentTest.ExpectedResponse); err != nil {
				newFailureMessage(currentTest.HTTPMethod, currentTest.URL(), currentTest.TestSuit, err.Error())
				return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
			}
		}
		// Call end function
		if currentTest.CustomEndFunc != nil {
			if err := currentTest.CustomEndFunc(res); err != nil {
				newFailureMessage(currentTest.HTTPMethod, currentTest.URL(), currentTest.TestSuit, err.Error())
				return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
			}
		}
		newSuccessMessage(currentTest.HTTPMethod, currentTest.URL(), currentTest.TestSuit)
	}
	return nil
}

func getBaseURL(environment string) (string, error) {
	switch environment {
	case "local":
		return "http://localhost:3000/", nil
	case "preprod":
		return "http://oldUser.preprod.alexandrio.cloud/", nil
	case "prod":
		return "http://oldUser.alexandrio.cloud/", nil
	default:
		return "", errors.New("provided environment unknown")
	}
}

func newSuccessMessage(verb string, route string, test string) {
	itgmtod.NewSuccessMessage(itgmtod.BackBlue("[USER]"), verb, route, test)
}

func newFailureMessage(verb string, route string, test string, message string) {
	itgmtod.NewFailureMessage(itgmtod.BackBlue("[USER]"), verb, route, test, message)
}
