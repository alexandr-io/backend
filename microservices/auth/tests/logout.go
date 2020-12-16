package tests

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func testLogoutWorking(baseURL string, jwt string) error {
	// Create a new request to auth route
	req, err := http.NewRequest(http.MethodPost, baseURL+"logout", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/logout", "Logout Suit", "Can't call "+baseURL+"logout")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusNoContent {
		newFailureMessage("POST", "/logout", "Logout Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusNoContent, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/logout", "Logout Suit")
	return nil
}

func testAlreadyLogout(baseURL string, jwt string) error {
	// Create a new request to auth route
	req, err := http.NewRequest(http.MethodPost, baseURL+"logout", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/logout", "Logout Suit", "Can't call "+baseURL+"logout")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		newFailureMessage("POST", "/logout", "Logout Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("POST", "/logout", "Logout Suit")
	return nil
}
