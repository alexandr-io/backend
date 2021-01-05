package tests

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var (
	green       = color.New(color.FgGreen).SprintFunc()
	red         = color.New(color.FgRed).SprintFunc()
	cyan        = color.New(color.FgCyan).SprintfFunc()
	magenta     = color.New(color.FgHiMagenta).SprintfFunc()
	backMagenta = color.New(color.BgMagenta).Add(color.FgWhite).SprintfFunc()
)

// ExecLibraryTests execute integration tests of library MS routes
func ExecLibraryTests(environment string, jwt string) error {
	rand.Seed(time.Now().UnixNano())
	var errorHappened = false

	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}

	if err = workingTestSuit(baseURL, jwt); err != nil {
		errorHappened = true
	}

	if errorHappened {
		return errors.New("error in library tests")
	}
	return nil
}

func getBaseURL(environment string) (string, error) {
	switch environment {
	case "local":
		return "http://localhost:3002/", nil
	case "preprod":
		return "http://library.preprod.alexandrio.cloud/", nil
	case "prod":
		return "http://library.alexandrio.cloud/", nil
	default:
		return "", errors.New("provided environment unknown")
	}
}

func workingTestSuit(baseURL string, jwt string) error {
	_, err := testLibraryCreateWorking(baseURL, jwt)
	if err != nil {
		return err
	}
	libraries, err := testLibrariesGetWorking(baseURL, jwt)
	if err != nil {
		return err
	}
	_, err = testLibraryGetWorking(baseURL, jwt, libraries.Libraries[0])
	if err != nil {
		return err
	}
	book, err := testBookCreateWorking(baseURL, jwt, libraries.Libraries[0])
	if err != nil {
		return err
	}
	_, err = testBookGetWorking(baseURL, jwt, libraries.Libraries[0], book)
	if err != nil {
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

	fmt.Printf("%s\t %s\t%-35s%-14s%-5s\n", backMagenta("[LIBRARY]"), coloredVerb, route, test, green("✓"))
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

	fmt.Printf("%s\t %s\t%-35s%-14s%-5s\t%s\n", backMagenta("[LIBRARY]"), coloredVerb, route, test, red("✗"), message)
}
