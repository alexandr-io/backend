package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testLibrariesGetWorking(baseURL string, jwt string) (*libraries, error) {
	// Create a new request to libraries get route
	req, err := http.NewRequest(http.MethodGet, baseURL+"libraries", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+jwt)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("GET", "/libraries", "Working Suit", "Can't call "+baseURL+"libraries")
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("GET", "/libraries", "Working Suit", fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return nil, errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		newFailureMessage("GET", "/libraries", "Working Suit", "Can't read response body")
		return nil, err
	}
	// Parse body data
	var bodyData libraries
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		newFailureMessage("GET", "/libraries", "Working Suit", "Can't unmarshal json")
		return nil, err
	}
	newSuccessMessage("GET", "/libraries", "Working Suit")
	return &bodyData, nil
}
