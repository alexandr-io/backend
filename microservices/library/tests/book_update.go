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

func testBookUpdateWorking(baseURL string, jwt string, libraryResponse libraryList, bookResponse *book) (*book, error) {
	// Create a new request to book post route
	payload := bytes.NewBuffer([]byte("{\"title\": \"The book was updated\"}"))
	req, err := http.NewRequest(http.MethodPost, baseURL+"library/"+libraryResponse.ID+"/book/"+bookResponse.ID, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("POST", "/library/:library_id/book/:book_id", "Working Suit", "Can't call "+baseURL+"/library/"+libraryResponse.ID+"/book/"+bookResponse.ID)
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("POST", "/library/:library_id/book/:book_id", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return nil, errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		newFailureMessage("POST", "/library/:library_id/book/:book_id", "Working Suit", "Can't read response body")
		return nil, err
	}
	// Parse body data
	var bodyData book
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		log.Println(body)
		newFailureMessage("POST", "/library/:library_id/book/:book_id", "Working Suit", "Can't unmarshal json")
		return nil, err
	}
	newSuccessMessage("POST", "/library/:library_id/book/:book_id", "Working Suit")
	return &bodyData, nil
}
