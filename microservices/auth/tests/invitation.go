package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testInvitationWorking(baseURL string, suit string) (*invitation, error) {
	// Create a new request to invitation new route
	req, err := http.NewRequest(http.MethodGet, JoinURL(baseURL, "/invitation/new"), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("GET", "/invitation/new", suit, "Can't call "+JoinURL(baseURL, "/invitation/new"))
		return nil, err
	}
	// Check returned http code
	if res.StatusCode != http.StatusOK {
		newFailureMessage("GET", "/invitation/new", suit, fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusOK, res.StatusCode))
		return nil, errors.New("")
	}
	// Read returned body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Parse body data
	var bodyData invitation
	if err := json.Unmarshal(body, &bodyData); err != nil {
		log.Println(err)
		return nil, err
	}
	newSuccessMessage("GET", "/invitation/new", suit)
	return &bodyData, nil
}

func testDeleteInvitationWorking(baseURL string, userData *user, inv invitation, suit string) error {
	// Create a new request to auth route
	req, err := http.NewRequest(http.MethodDelete, JoinURL(baseURL, "/invitation/"+inv.Token), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+userData.AuthToken)
	// Exec request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		newFailureMessage("DELETE", "/invitation/"+inv.Token, suit, "Can't call "+JoinURL(baseURL, "/invitation/"+inv.Token))
		return err
	}
	// Check returned http code
	if res.StatusCode != http.StatusNoContent {
		newFailureMessage("DELETE", "/invitation/"+inv.Token, suit, fmt.Sprintf("[Expected: %d,\tGot: %d]", http.StatusNoContent, res.StatusCode))
		return errors.New("")
	}

	newSuccessMessage("DELETE", "/invitation/"+inv.Token, suit)
	return nil
}
