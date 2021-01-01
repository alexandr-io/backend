package tests

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func testResetPasswordWorking(baseURL string, userData user) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"email\": \"" + userData.Email + "\"}"))
	// Create a new request to password reset route
	req, err := http.NewRequest(http.MethodPost, JoinURL(baseURL, "password/reset"), payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/password/reset", "Working Suit", "Can't call "+JoinURL(baseURL, "password/reset"))
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusNoContent {
		newFailureMessage("POST", "/password/reset", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusNoContent, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/password/reset", "Working Suit")
	return nil
}

func testResetPasswordBadRequest(baseURL string) error {
	// Create an incorrect payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"email\": \"wrong-email\"}"))
	// Create a new request to login route
	req, err := http.NewRequest(http.MethodPost, JoinURL(baseURL, "password/reset"), payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/password/reset", "Bad Request", "Can't call "+JoinURL(baseURL, "password/reset"))
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("POST", "/password/reset", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/password/reset", "Bad Request")
	return nil
}

func testResetPasswordNoMatch(baseURL string) error {
	// Generate random string for username and email usage
	randomName := randStringRunes(12)
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"email\": \"" + randomName + "@test.test\"}"))
	// Create a new request to login route
	req, err := http.NewRequest(http.MethodPost, JoinURL(baseURL, "password/reset"), payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/password/reset", "No Match", "Can't call "+JoinURL(baseURL, "password/reset"))
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		newFailureMessage("POST", "/password/reset", "No Match", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/password/reset", "No Match")
	return nil
}
