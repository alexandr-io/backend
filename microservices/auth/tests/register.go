package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randStringRunes generate a random string of specified length
func randStringRunes(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return "test-" + string(b)
}

func testRegisterWorking(baseURL string) (*user, error) {
	// Generate random string for username and email usage
	randomName := randStringRunes(12)
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"username\": \"" + randomName + "\", \"email\": \"" + randomName + "@test.com\", \"password\": \"test\", \"confirm_password\": \"test\"}"))
	// Create a new request to register route
	req, err := http.NewRequest(http.MethodPost, baseURL+"register", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/register", "Working Suit", "Can't call "+baseURL+"register")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusCreated {
		newFailureMessage("POST", "/register", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusCreated, res.StatusCode))
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
	newSuccessMessage("POST", "/register", "Working Suit")
	return &bodyData, nil
}

func testRegisterBadRequest(baseURL string) error {
	// Create an incorrect payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"userna\": \"random\", \"password\": \"test\"}"))
	// Create a new request to register route
	req, err := http.NewRequest(http.MethodPost, baseURL+"register", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/register", "Bad Request", "Can't call "+baseURL+"register")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("POST", "/register", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/register", "Bad Request")
	return nil
}

func testRegisterDuplicate(baseURL string, userData user) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"username\": \"" + userData.Username + "\", \"email\": \"" + userData.Email + "\", \"password\": \"test\", \"confirm_password\": \"test\"}"))
	// Create a new request to register route
	req, err := http.NewRequest(http.MethodPost, baseURL+"register", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/register", "Duplicate", "Can't call "+baseURL+"register")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("POST", "/register", "Duplicate", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/register", "Duplicate")
	return nil
}
