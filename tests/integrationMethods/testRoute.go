package integrationMethods

import (
	"fmt"
	"io"
	"net/http"
)

// TestRoute create and call an HTTP route with the given parameters. In case of error, an error is returned.
// The response object is returned in case of success.
func TestRoute(HTTPMethod string, URL string, authJWT *string, body io.Reader, expectedHTTPCode int) (*http.Response, error) {
	// Create request with body if given
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(HTTPMethod, URL, body)
	} else {
		req, err = http.NewRequest(HTTPMethod, URL, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Add Bearer authentication token if given
	if authJWT != nil {
		req.Header.Set("Authorization", "Bearer "+*authJWT)
	}

	// Execute request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check returned http code
	if res.StatusCode != expectedHTTPCode {
		return nil, fmt.Errorf("[Expected: %d,\tGot: %d]", expectedHTTPCode, res.StatusCode)
	}
	return res, nil
}
