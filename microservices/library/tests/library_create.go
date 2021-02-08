package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexandr-io/backend/library/data"
)

// LibrayCreateEndFunction is a function called at the end of a library create test
func LibrayCreateEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var libraryData data.Library
	if err := json.Unmarshal(resBody, &libraryData); err != nil {
		return err
	}
	libraryID = libraryData.ID
	return nil
}
