package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexandr-io/backend/auth/data"
)

func invitationEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var invitationData data.Invitation
	if err := json.Unmarshal(resBody, &invitationData); err != nil {
		return err
	}
	// Store invitation token
	invitationToken = invitationData.Token
	return nil
}
