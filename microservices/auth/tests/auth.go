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
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+userData.AuthToken)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[AUTH]: GET\t/\t\tWorking Suit\t✗\tCan't call " + baseURL)
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		errorString := fmt.Sprintf("[AUTH]: GET\t/\t\tWorking Suit\t✗\t[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	// Compare body with expected result
	if string(body) != expectedResult {
		errorString := fmt.Sprintf("[AUTH]: GET\t/\t\tWorking Suit\t✗\t[Expected: %s,\tGot: %s]", expectedResult, string(body))
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	fmt.Println("[AUTH]: GET\t/\t\tWorking Suit\t✓")
	return nil
}

func testAuthInvalidToken(baseURL string) error {
	// Create a new request to auth route
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer randomString")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[AUTH]: GET\t/\t\tInvalid Token\t✗\tCan't call " + baseURL)
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		errorString := fmt.Sprintf("[AUTH]: GET\t/\t\tInvalid Token\t✗\t[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}

	fmt.Println("[AUTH]: GET\t/\t\tInvalid Token\t✓")
	return nil
}
