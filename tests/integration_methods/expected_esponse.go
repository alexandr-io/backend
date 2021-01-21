package integration_methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// CheckExpectedResponse will read and unmarshall the response body and then call StructContain
func CheckExpectedResponse(res *http.Response, expected interface{}) error {
	// Read response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body = ioutil.NopCloser(bytes.NewReader(resBody))

	// Get expected type and create new interface to unmarshall response body
	expectedType := reflect.TypeOf(expected)
	responseDataValue := reflect.New(expectedType)
	responseData := responseDataValue.Interface()

	// Parse response body
	if err := json.Unmarshal(resBody, responseData); err != nil {
		return err
	}

	if !StructContain(responseData, expected) {
		return fmt.Errorf("response struct does not correspond to expected response")
	}
	return nil
}
