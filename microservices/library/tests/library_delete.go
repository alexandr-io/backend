package tests

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func testLibraryDeleteWorking(baseURL string, jwt string, libraryResponse libraryList) error {
	// Create a new request to library delete route
	payload := bytes.NewBuffer([]byte("{\"name\": \"" + libraryResponse.Name + "\"}"))
	req, err := http.NewRequest(http.MethodDelete, baseURL+"library", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("DELETE", "/library", "Working Suit", "Can't call "+baseURL+"library")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusNoContent {
		newFailureMessage("DELETE", "/library", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusNoContent, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("DELETE", "/library", "Working Suit")
	return nil
}

func testLibraryDeleteBadRequest(baseURL string, jwt string) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"name\": 42}"))
	// Create a new request to user update route
	req, err := http.NewRequest(http.MethodDelete, baseURL+"library", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("DELETE", "/library", "Bad Request", "Can't call "+baseURL+"library")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("DELETE", "/library", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("DELETE", "/library", "Bad Request")
	return nil
}
