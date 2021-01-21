package tests

import (
	"encoding/json"
	"github.com/alexandr-io/backend/auth/data"
	"io/ioutil"
	"net/http"
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
