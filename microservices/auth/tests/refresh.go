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
		fmt.Println("[AUTH]: POST\t/refresh\tWorking Suit\t✗\tCan't call " + baseURL + "refresh")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		errorString := fmt.Sprintf("[AUTH]: POST\t/refresh\tWorking Suit\t✗\t[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode)
		fmt.Println(errorString)
		return nil, errors.New(errorString)
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
	fmt.Println("[AUTH]: POST\t/refresh\tWorking Suit\t✓")
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
		fmt.Println("[AUTH]: POST\t/refresh\tInvalid Token\t✗\tCan't call " + baseURL + "refresh")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusUnauthorized {
		errorString := fmt.Sprintf("[AUTH]: POST\t/refresh\tInvalid Token\t✗\t[Expected: %d,\tGot: %d]", http.StatusUnauthorized, res.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	fmt.Println("[AUTH]: POST\t/refresh\tInvalid Token\t✓")
	return nil
}
