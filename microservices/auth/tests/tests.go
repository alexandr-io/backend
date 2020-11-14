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

	err = workingTestSuit(baseURL)
	if err != nil {
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
