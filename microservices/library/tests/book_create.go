package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// BookCreateEndFunction is a function called at the end of a book creation test
func BookCreateEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	bookData := struct {
		ID string `json:"id"`
	}{}
	if err = json.Unmarshal(resBody, &bookData); err != nil {
		return err
	}
	bookID = bookData.ID
	return nil
}
