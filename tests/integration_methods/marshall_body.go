package integration_methods

import "encoding/json"

// MarshallBody is used to get the JSON bytes out of a struct
func MarshallBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}
