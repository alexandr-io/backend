package tests

import (
	"bytes"
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
	return string(b)
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
		fmt.Println("[AUTH]: POST\t/register\t-> Fail\tCan't call " + baseURL + "register")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusCreated {
		errorString := fmt.Sprintf("[AUTH]: POST\t/register\t-> Fail\t[Expected: %d,\tGot: %d]", http.StatusCreated, res.StatusCode)
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
	fmt.Println("[AUTH]: POST\t/register\t-> Success")
	return &bodyData, nil
}
