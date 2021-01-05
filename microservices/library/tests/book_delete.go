package tests

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func testBookDeleteWorking(baseURL string, jwt string, libraryResponse libraryList, bookResponse *book) error {
	// Create a new request to book delete route
	payload := bytes.NewBuffer([]byte("{\"book_id\": \"" + bookResponse.ID + "\", \"library_id\": \"" + libraryResponse.ID + "\"}"))
	req, err := http.NewRequest(http.MethodDelete, baseURL+"book", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("DELETE", "/book", "Working Suit", "Can't call "+baseURL+"book")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusNoContent {
		newFailureMessage("DELETE", "/book", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusNoContent, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("DELETE", "/book", "Working Suit")
	return nil
}

func testBookDeleteBadRequest(baseURL string, jwt string) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"book_id\": 42, \"library_id\": 42}"))
	// Create a new request to user update route
	req, err := http.NewRequest(http.MethodDelete, baseURL+"book", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("DELETE", "/book", "Bad Request", "Can't call "+baseURL+"book")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("DELETE", "/book", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("DELETE", "/book", "Bad Request")
	return nil
}
