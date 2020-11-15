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

func testLoginWorking(baseURL string, userData *user) (*user, error) {
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
		fmt.Println("[AUTH]: POST\t/login\t\tWorking Suit\t✗\tCan't call " + baseURL + "login")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		errorString := fmt.Sprintf("[AUTH]: POST\t/login\t\tWorking Suit\t✗\t[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode)
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
	fmt.Println("[AUTH]: POST\t/login\t\tWorking Suit\t✓")
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
		fmt.Println("[AUTH]: POST\t/login\t\tBad Request\t✗\tCan't call " + baseURL + "login")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		errorString := fmt.Sprintf("[AUTH]: POST\t/login\t\tBad Request\t✗\t[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	fmt.Println("[AUTH]: POST\t/login\t\tBad Request\t✓")
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
		fmt.Println("[AUTH]: POST\t/login\t\tNo Match\t✗\tCan't call " + baseURL + "login")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		errorString := fmt.Sprintf("[AUTH]: POST\t/login\t\tNo Match\t✗\t[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	fmt.Println("[AUTH]: POST\t/login\t\tNo Match\t✓")
	return nil
}
