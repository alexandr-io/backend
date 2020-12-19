package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testRefreshWorking(baseURL string, userData *user) (*user, error) {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"refresh_token\": \"" + userData.RefreshToken + "\"}"))
	// Create a new request to refresh route
	req, err := http.NewRequest(http.MethodPost, baseURL+"refresh", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/refresh", "Working Suit", "Can't call "+baseURL+"refresh")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("POST", "/refresh", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return nil, errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Parse body data
	var bodyData user
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		return nil, err
	}
	newSuccessMessage("POST", "/refresh", "Working Suit")
	return &bodyData, nil
}

func testRefreshInvalidToken(baseURL string) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"refresh_token\": \"randomString\"}"))
	// Create a new request to refresh route
	req, err := http.NewRequest(http.MethodPost, baseURL+"refresh", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/refresh", "Invalid Token", "Can't call "+baseURL+"refresh")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		newFailureMessage("POST", "/refresh", "Invalid Token", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/refresh", "Invalid Token")
	return nil
}
