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
		fmt.Println("[AUTH]: POST\t/refresh\t-> Fail\tCan't call " + baseURL + "refresh")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		errorString := fmt.Sprintf("[AUTH]: POST\t/refresh\t-> Fail\t[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode)
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
	fmt.Println("[AUTH]: POST\t/refresh\t-> Success")
	return &bodyData, nil
}
