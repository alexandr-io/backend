package tests

import (
	"bytes"
	"context"
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
		fmt.Println("[AUTH]: POST\t/register\tWorking Suit\t✗\tCan't call " + baseURL + "register")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusCreated {
		errorString := fmt.Sprintf("[AUTH]: POST\t/register\tWorking Suit\t✗\t[Expected: %d,\tGot: %d]", http.StatusCreated, res.StatusCode)
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
	fmt.Println("[AUTH]: POST\t/register\tWorking Suit\t✓")
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
		fmt.Println("[AUTH]: POST\t/register\tBad Request\t✗\tCan't call " + baseURL + "register")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusBadRequest {
		errorString := fmt.Sprintf("[AUTH]: POST\t/register\tBad Request\t✗\t[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	fmt.Println("[AUTH]: POST\t/register\tBad Request\t✓")
	return nil
}

func testRegisterDuplicate(baseURL string) error {
	// Generate random string for username and email usage
	randomName := randStringRunes(12)
	// Create payload to send to the route
	payload := bytes.NewBuffer([]byte("{\"username\": \"" + randomName + "\", \"email\": \"" + randomName + "@test.com\", \"password\": \"test\", \"confirm_password\": \"test\"}"))
	// Create a new request to register route
	req, err := http.NewRequest(http.MethodPost, baseURL+"register", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	clone := req.Clone(context.TODO())
	clone.Body, err = req.GetBody()
	if err != nil {
		log.Println(err)
		return err
	}
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[AUTH]: POST\t/register\tDuplicate\t✗\tCan't call " + baseURL + "register")
		return err
	}
	cloneRes, err := http.DefaultClient.Do(clone)
	if err != nil {
		fmt.Println("[AUTH]: POST\t/register\tDuplicate\t✗\tCan't call " + baseURL + "register")
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusCreated {
		errorString := fmt.Sprintf("[AUTH]: POST\t/register\tDuplicate\t✗\t[Expected: %d,\tGot: %d]", http.StatusBadRequest, res.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	if cloneRes.StatusCode != http.StatusBadRequest {
		errorString := fmt.Sprintf("[AUTH]: POST\t/register\tDuplicate\t✗\t[Expected: %d,\tGot: %d]", http.StatusBadRequest, cloneRes.StatusCode)
		fmt.Println(errorString)
		return errors.New(errorString)
	}
	fmt.Println("[AUTH]: POST\t/register\tDuplicate\t✓")
	return nil
}
