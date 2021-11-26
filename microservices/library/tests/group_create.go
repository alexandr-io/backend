package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexandr-io/backend/library/data/permissions"
)

// GroupPostEndFunction retrieve the result of a group creation and store the group's ID
func GroupPostEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var groupData permissions.Group
	if err := json.Unmarshal(resBody, &groupData); err != nil {
		return err
	}
	groupID = groupData.ID.Hex()
	return nil
}
