package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexandr-io/backend/library/data"
)

// UserDataCreateEndFunction is a function called at the end of a user data create test
func UserDataCreateEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var userData data.UserData
	if err = json.Unmarshal(resBody, &userData); err != nil {
		return err
	}
	dataID = userData.ID.Hex()
	return nil
}
