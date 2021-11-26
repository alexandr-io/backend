package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexandr-io/backend/library/data"
)

// LibrariesGetEndFunction is a function called at the end of a library get test
func LibrariesGetEndFunction(res *http.Response) error {
	var librariesData []data.Library
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	if err = json.Unmarshal(resBody, &librariesData); err != nil {
		return err
	}

	for _, library := range librariesData {
		if library.Name == libraryName {
			libraryID = library.ID
			break
		}
	}
	return nil
}
