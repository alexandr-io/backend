// Package handlers Documentation of Library API
//
// Documentation for the Library microservice REST API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/berrors"
)

//
// Input
//

//
// Output
//

// A library
// swagger:response libraryResponse
type libraryResponseWrapper struct {
	// A single library
	// in: body
	Body data.Library
}

// A list of the user libraries names
// swagger:response librariesNamesResponse
type librariesNamesWrapper struct {
	// The list of the names of the libraries
	// in: body
	Body data.LibrariesNames
}

//
// Error
//

// A bad request error response
// swagger:response badRequestErrorResponse
type badRequestErrorResponseWrapper struct {
	// A list of described bad request error
	// in: body
	Body berrors.BadInput
}
