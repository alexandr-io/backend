package tests

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var (
	green    = color.New(color.FgGreen).SprintFunc()
	red      = color.New(color.FgRed).SprintFunc()
	cyan     = color.New(color.FgCyan).SprintfFunc()
	magenta  = color.New(color.FgHiMagenta).SprintfFunc()
	backCyan = color.New(color.BgCyan).Add(color.FgBlack).SprintfFunc()
)

// ExecAuthTests execute integration tests of auth MS routes
func ExecAuthTests(environment string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var errorHappened = false

	baseURL, err := getBaseURL(environment)
	if err != nil {
		return "", err
	}

	jwt, err := workingTestSuit(baseURL)
	if err != nil {
		errorHappened = true
	}
	if err = badRequestTests(baseURL); err != nil {
		errorHappened = true
	}
	if err = incorrectTests(baseURL); err != nil {
		errorHappened = true
	}

	if errorHappened {
		return "", errors.New("error in auth tests")
	}
	return jwt, nil
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

func workingTestSuit(baseURL string) (string, error) {
	userData, err := testRegisterWorking(baseURL)
	if err != nil {
		return "", err
	}
	if err := testAuthWorking(baseURL, userData); err != nil {
		return "", err
	}

	userDataLogin, err := testLoginWorking(baseURL, userData)
	if err != nil {
		return "", err
	}
	if err := testAuthWorking(baseURL, userDataLogin); err != nil {
		return "", err
	}

	userDataRefresh, err := testRefreshWorking(baseURL, userDataLogin)
	if err != nil {
		return "", err
	}
	if err := testAuthWorking(baseURL, userDataRefresh); err != nil {
		return "", err
	}

	return userData.AuthToken, nil
}

func badRequestTests(baseURL string) error {
	if err := testRegisterBadRequest(baseURL); err != nil {
		return err
	}
	if err := testLoginBadRequest(baseURL); err != nil {
		return err
	}
	return nil
}

func incorrectTests(baseURL string) error {
	if err := testRegisterDuplicate(baseURL); err != nil {
		return err
	}
	if err := testLoginNoMatch(baseURL); err != nil {
		return err
	}
	if err := testAuthInvalidToken(baseURL); err != nil {
		return err
	}
	if err := testRefreshInvalidToken(baseURL); err != nil {
		return err
	}
	return nil
}

func newSuccessMessage(verb string, route string, test string) {
	coloredVerb := verb
	switch verb {
	case "GET":
		coloredVerb = green(verb)
		break
	case "POST":
		coloredVerb = cyan(verb)
		break
	case "PUT":
		coloredVerb = magenta(verb)
		break
	case "DELETE":
		coloredVerb = red(verb)
		break
	}

	fmt.Printf("%s\t %s\t%-20s%-14s%-5s\n", backCyan("[AUTH]"), coloredVerb, route, test, green("✓"))
}

func newFailureMessage(verb string, route string, test string, message string) {
	coloredVerb := verb
	switch verb {
	case "GET":
		coloredVerb = green(verb)
		break
	case "POST":
		coloredVerb = cyan(verb)
		break
	case "PUT":
		coloredVerb = magenta(verb)
		break
	case "DELETE":
		coloredVerb = red(verb)
		break
	}

	fmt.Printf("%s\t %s\t%-20s%-14s%-5s\t%s\n", backCyan("[AUTH]"), coloredVerb, route, test, red("✗"), message)
}
