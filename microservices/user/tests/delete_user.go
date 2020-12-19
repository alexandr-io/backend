package tests

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func testUserDeleteWorking(baseURL string, jwt string) error {
	// Create a new request to user delete route
	req, err := http.NewRequest(http.MethodDelete, baseURL+"user", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("DELETE", "/user", "Working Suit", "Can't call "+baseURL+"user")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusNoContent {
		newFailureMessage("DELETE", "/user", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusNoContent, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("DELETE", "/user", "Working Suit")
	return nil
}

func testUserAlreadyDelete(baseURL string, jwt string) error {
	// Create a new request to user delete route
	req, err := http.NewRequest(http.MethodDelete, baseURL+"user", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("DELETE", "/user", "Bad Request", "Can't call "+baseURL+"user")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		newFailureMessage("DELETE", "/user", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("DELETE", "/user", "Bad Request")
	return nil
}
