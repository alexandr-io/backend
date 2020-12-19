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

func testLoginWorking(baseURL string, userData user) (*user, error) {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"login\": \"" + userData.Email + "\", \"password\": \"test\"}"))
	// Create a new request to login route
	req, err := http.NewRequest(http.MethodPost, baseURL+"login", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/login", "Working Suit", "Can't call "+baseURL+"login")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("POST", "/login", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
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
	newSuccessMessage("POST", "/login", "Working Suit")
	return &bodyData, nil
}

func testLoginBadRequest(baseURL string) error {
	// Create an incorrect payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"logi\": \"random\"}"))
	// Create a new request to login route
	req, err := http.NewRequest(http.MethodPost, baseURL+"login", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/login", "Bad Request", "Can't call "+baseURL+"login")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("POST", "/login", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/login", "Bad Request")
	return nil
}

func testLoginNoMatch(baseURL string) error {
	// Generate random string for username and email usage
	randomName := randStringRunes(12)
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"login\": \"" + randomName + "\", \"password\": \"" + randomName + "\"}"))
	// Create a new request to login route
	req, err := http.NewRequest(http.MethodPost, baseURL+"login", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/login", "No Match", "Can't call "+baseURL+"login")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("POST", "/login", "No Match", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/login", "No Match")
	return nil
}
