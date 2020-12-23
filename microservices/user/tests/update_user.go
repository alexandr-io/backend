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

func testUserUpdateWorking(baseURL string, userData *user) error {
	newUsername := userData.Username[:len(userData.Username)-1]
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"username\": \"" + newUsername + "\"}"))
	// Create a new request to user update route
	req, err := http.NewRequest(http.MethodPut, baseURL+"user", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userData.AuthToken)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("PUT", "/user", "Working Suit", "Can't call "+baseURL+"user")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("PUT", "/user", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	// Parse body data
	var bodyData user
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		return err
	}
	if bodyData.Username != newUsername &&
		bodyData.Email != userData.Email {
		newFailureMessage("PUT", "/user", "Working Suit", "Data != expected data")
		return errors.New("")
	}
	newSuccessMessage("PUT", "/user", "Working Suit")
	return nil
}

func testUserUpdateBadRequest(baseURL string, jwt string) error {
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"username\": 42}"))
	// Create a new request to user update route
	req, err := http.NewRequest(http.MethodPut, baseURL+"user", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("PUT", "/user", "Bad Request", "Can't call "+baseURL+"user")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		newFailureMessage("PUT", "/user", "Bad Request", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode))
		return errors.New("")
	}
	newSuccessMessage("PUT", "/user", "Bad Request")
	return nil
}
