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

func testBookCreateWorking(baseURL string, jwt string, libraryResponse libraryList) (*book, error) {
	// Create a new request to book post route
	payload := bytes.NewBuffer([]byte("{\"title\": \"The book\", \"author\": \"The author\", \"description\": \"The description\", \"tags\": [\"The 1st tag\", \"the 2nd tag\"], \"library_id\": \"" + libraryResponse.ID + "\"}"))
	req, err := http.NewRequest(http.MethodPost, baseURL+"book", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/book", "Working Suit", "Can't call "+baseURL+"book")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusCreated {
		newFailureMessage("POST", "/book", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusCreated, res.StatusCode))
		return nil, errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		newFailureMessage("POST", "/book", "Working Suit", "Can't read response body")
		return nil, err
	}
	// Parse body data
	var bodyData book
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		newFailureMessage("POST", "/book", "Working Suit", "Can't unmarshal json")
		return nil, err
	}
	newSuccessMessage("POST", "/book", "Working Suit")
	return &bodyData, nil
}

func testBookCreateBadRequest(baseURL string, jwt string) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"title\": 42, \"library_id\": \"library_does_not_exist\"}"))
	// Create a new request to user update route
	req, err := http.NewRequest(http.MethodPost, baseURL+"book", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/book", "Bad Request", "Can't call "+baseURL+"book")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("POST", "/book", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/book", "Bad Request")
	return nil
}
