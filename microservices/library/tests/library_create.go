package tests

import (
	"encoding/json"
	"github.com/alexandr-io/backend/library/data"
	"io/ioutil"
	"net/http"
)

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