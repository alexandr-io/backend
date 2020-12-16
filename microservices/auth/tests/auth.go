package tests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testAuthWorking(baseURL string, userData *user) error {
	// Describe expected result
	expectedResult := fmt.Sprintf("{\"username\":\"%s\"}", userData.Username)
	// Create a new request to auth route
	req, err := http.NewRequest(http.MethodGet, baseURL+"auth", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+userData.AuthToken)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("GET", "/auth", "Working Suit", "Can't call "+baseURL+"auth")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("GET", "/auth", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	// Compare body with expected result
	if string(body) != expectedResult {
		newFailureMessage("GET", "/auth", "Working Suit", fmt.Sprintf("[Expected: %s,\tGot: %s]", expectedResult, string(body)))
		return errors.New("")
	}
	newSuccessMessage("GET", "/auth", "Working Suit")
	return nil
}

func testAuthInvalidToken(baseURL string) error {
	// Create a new request to auth route
	req, err := http.NewRequest(http.MethodGet, baseURL+"auth", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer randomString")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("GET", "/auth", "Invalid Token", "Can't call "+baseURL+"auth")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		newFailureMessage("GET", "/auth", "Invalid Token", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode))
		return errors.New("")
	}

	newSuccessMessage("GET", "/auth", "Invalid Token")
	return nil
}

func testAuthLogoutToken(baseURL string, userData *user) error {
	// Create a new request to auth route
	req, err := http.NewRequest(http.MethodGet, baseURL+"auth", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+userData.AuthToken)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("GET", "/auth", "Logout Token", "Can't call "+baseURL+"auth")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		newFailureMessage("GET", "/auth", "Logout Token", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode))
		return errors.New("")
	}

	newSuccessMessage("GET", "/auth", "Logout Token")
	return nil
}
