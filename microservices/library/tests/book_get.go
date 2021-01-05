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

func testBookGetWorking(baseURL string, jwt string, libraryResponse libraryList, bookResponse *book) (*book, error) {
	// Create a new request to book get route
	payload := bytes.NewBuffer([]byte("{\"book_id\": \"" + bookResponse.ID + "\", \"library_id\": \"" + libraryResponse.ID + "\"}"))
	req, err := http.NewRequest(http.MethodPut, baseURL+"book", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("PUT", "/book", "Working Suit", "Can't call "+baseURL+"user")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("PUT", "/book", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return nil, errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		newFailureMessage("PUT", "/book", "Working Suit", "Can't read response body")
		return nil, err
	}
	// Parse body data
	var bodyData book
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		newFailureMessage("PUT", "/book", "Working Suit", "Can't unmarshal json")
		return nil, err
	}
	newSuccessMessage("PUT", "/book", "Working Suit")
	return &bodyData, nil
}
