package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func refreshEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var refreshData user
	if err := json.Unmarshal(resBody, &refreshData); err != nil {
		return err
	}
	// Store tokens for future uses
	authTokenRefresh = refreshData.AuthToken
	return nil
}
