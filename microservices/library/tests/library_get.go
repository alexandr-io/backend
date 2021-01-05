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

func testLibraryGetWorking(baseURL string, jwt string, libraryResponse libraryList) (*library, error) {
	// Create a new request to library put route
	payload := bytes.NewBuffer([]byte("{\"name\": \"" + libraryResponse.Name + "\"}"))
	req, err := http.NewRequest(http.MethodPut, baseURL+"library", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("PUT", "/library", "Working Suit", "Can't call "+baseURL+"library")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("PUT", "/library", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return nil, errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		newFailureMessage("PUT", "/library", "Working Suit", "Can't read response body")
		return nil, err
	}
	// Parse body data
	var bodyData library
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		newFailureMessage("PUT", "/library", "Working Suit", "Can't unmarshal json")
		return nil, err
	}
	newSuccessMessage("PUT", "/library", "Working Suit")
	return &bodyData, nil
}

func testLibraryRetrieveBadRequest(baseURL string, jwt string) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"name\": 42}"))
	// Create a new request to user update route
	req, err := http.NewRequest(http.MethodPut, baseURL+"library", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("PUT", "/library", "Bad Request", "Can't call "+baseURL+"library")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("PUT", "/library", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("PUT", "/library", "Bad Request")
	return nil
}
