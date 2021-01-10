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
	backBlue = color.New(color.BgBlue).Add(color.FgWhite).SprintfFunc()
)

// ExecUserTests execute integration tests of user MS routes
func ExecUserTests(environment string, jwt string) error {
	rand.Seed(time.Now().UnixNano())
	var errorHappened = false

	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}

	if err = workingTestSuit(baseURL, jwt); err != nil {
		errorHappened = true
	}
	if err = badInputTestSuit(baseURL, jwt); err != nil {
		errorHappened = true
	}
	if err = deleteUser(baseURL, jwt); err != nil {
		errorHappened = true
	}

	if errorHappened {
		return errors.New("error in user tests")
	}
	return nil
}

func getBaseURL(environment string) (string, error) {
	switch environment {
	case "local":
		return "http://localhost:3000/", nil
	case "preprod":
		return "http://user.preprod.alexandrio.cloud/", nil
	case "prod":
		return "http://user.alexandrio.cloud/", nil
	default:
		return "", errors.New("provided environment unknown")
	}
}

func workingTestSuit(baseURL string, jwt string) error {
	userData, err := testUserGetWorking(baseURL, jwt)
	if err != nil {
		return err
	}
	userData.AuthToken = jwt
	if err := testUserUpdateWorking(baseURL, userData); err != nil {
		return err
	}
	return nil
}

func badInputTestSuit(baseURL string, jwt string) error {
	if err := testUserUpdateBadRequest(baseURL, jwt); err != nil {
		return err
	}
	return nil
}

func deleteUser(baseURL string, jwt string) error {
	if err := testUserDeleteWorking(baseURL, jwt); err != nil {
		return err
	}
	if err := testUserAlreadyDelete(baseURL, jwt); err != nil {
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

	fmt.Printf("%-23s%-17s%-35s%-14s%-5s\n", backBlue("[USER]"), coloredVerb, route, test, green("✓"))
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

	fmt.Printf("%-23s%-17s%-35s%-14s%-5s\t%s\n", backBlue("[USER]"), coloredVerb, route, test, red("✗"), message)
}
