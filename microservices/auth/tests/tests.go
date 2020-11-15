package tests

import (
	"errors"
	"math/rand"
	"time"
)

// ExecAuthTests execute integration tests of auth MS routes
func ExecAuthTests(environment string) error {
	rand.Seed(time.Now().UnixNano())
	var errorHappened = false

	baseURL, err := getBaseURL(environment)
	if err != nil {
		return err
	}

	if err = workingTestSuit(baseURL); err != nil {
		errorHappened = true
	}
	if err = badRequestTests(baseURL); err != nil {
		errorHappened = true
	}
	if err = incorrectTests(baseURL); err != nil {
		errorHappened = true
	}

	if errorHappened {
		return errors.New("error in auth tests")
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
		return "http://auth.alexandrio.cloud", nil
	default:
		return "", errors.New("provided environment unknown")
	}
}

func workingTestSuit(baseURL string) error {
	userData, err := testRegisterWorking(baseURL)
	if err != nil {
		return err
	}
	if err := testAuthWorking(baseURL, userData); err != nil {
		return err
	}

	userDataLogin, err := testLoginWorking(baseURL, userData)
	if err != nil {
		return err
	}
	if err := testAuthWorking(baseURL, userDataLogin); err != nil {
		return err
	}

	userDataRefresh, err := testRefreshWorking(baseURL, userDataLogin)
	if err != nil {
		return err
	}
	if err := testAuthWorking(baseURL, userDataRefresh); err != nil {
		return err
	}

	return nil
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
