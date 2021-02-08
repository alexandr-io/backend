package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexandr-io/backend/library/data"
)

// bookCreation defines the structure of an API book for creation. copy for test
type bookCreation struct {
	Title       string   `json:"title,omitempty"`
	Author      string   `json:"author,omitempty"`
	Publisher   string   `json:"publisher,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	LibraryID   *string  `json:"library_id,omitempty"`
}

// BookCreateEndFunction is a function called at the end of a book creation test
func BookCreateEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var bookData data.Book
	if err := json.Unmarshal(resBody, &bookData); err != nil {
		return err
	}
	bookID = bookData.ID
	return nil
}
