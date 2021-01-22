package tests

import (
	"encoding/json"
	"fmt"
	"github.com/alexandr-io/backend/library/data"
	"io/ioutil"
	"net/http"
)

func LibrariesGetEndFunction(res *http.Response) error {
	// Read response Body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Parse response Body
	var librariesData data.LibrariesNames
	if err := json.Unmarshal(resBody, &librariesData); err != nil {
		return err
	}
	if librariesData.Libraries[0].Name != libraryName {
		return fmt.Errorf("expected: %s\t got: %s", libraryName, librariesData.Libraries[0].Name)
	}
	libraryID = librariesData.Libraries[0].ID
	return nil
}
