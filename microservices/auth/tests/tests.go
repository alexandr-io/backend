package tests

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/alexandr-io/backend/tests/integrationMethods"
	"net/http"
)

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
		url := integrationMethods.JoinURL(baseURL, currentTest.URL())
		body, err := integrationMethods.MarshallBody(currentTest.Body)
		if err != nil {
			newFailureMessage(currentTest.HTTPMethod, url, currentTest.TestSuit, err.Error())
			return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
		}

		// Test the route
		res, err := integrationMethods.TestRoute(
			currentTest.HTTPMethod,
			url,
			currentTest.AuthJWT,
			bytes.NewReader(body),
			currentTest.ExpectedHTTPCode)
		if err != nil {
			newFailureMessage(currentTest.HTTPMethod, url, currentTest.TestSuit, err.Error())
			return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
		}
		// Check expected response Body
		if currentTest.ExpectedResponse != nil {
			if err := integrationMethods.CheckExpectedResponse(res, currentTest.ExpectedResponse); err != nil {
				newFailureMessage(currentTest.HTTPMethod, url, currentTest.TestSuit, err.Error())
				return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
			}
		}
		// Call end function
		if currentTest.CustomEndFunc != nil {
			if err := currentTest.CustomEndFunc(res); err != nil {
				newFailureMessage(currentTest.HTTPMethod, url, currentTest.TestSuit, err.Error())
				return fmt.Errorf("error in " + currentTest.TestSuit + " test suit")
			}
		}
		newSuccessMessage(currentTest.HTTPMethod, url, currentTest.TestSuit)
	}
	return nil
}

func getBaseURL(environment string) (string, error) {
	switch environment {
	case "local":
		return "http://localhost:3001/", nil
	case "preprod":
		return "http://auth.preprod.alexandrio.cloud/", nil
	case "prod":
		return "http://auth.alexandrio.cloud/", nil
	default:
		return "", errors.New("provided environment unknown")
	}
}

func newSuccessMessage(verb string, route string, test string) {
	integrationMethods.NewSuccessMessage(integrationMethods.BackCyan("[AUTH]"), verb, route, test)
}

func newFailureMessage(verb string, route string, test string, message string) {
	integrationMethods.NewFailureMessage(integrationMethods.BackCyan("[AUTH]"), verb, route, test, message)
}
