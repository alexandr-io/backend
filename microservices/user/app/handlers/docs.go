// Package handlers Documentation of User API
//
// Documentation for the User microservice REST API
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
	"github.com/alexandr-io/berrors"
)

// A bad request error response
// swagger:response badRequestErrorResponse
type badRequestErrorResponseWrapper struct {
	// A list of described bad request error
	// in: body
	Body berrors.BadInput
}

// An unauthorized error response
// swagger:response unauthorizedErrorResponse
type unauthorizedErrorResponseWrapper struct {
	// The description of the unauthorized error
	// in: body
	Body struct {
		// The error message
		// Required: true
		// Example: Invalid or expired JWT
		Error string `json:"error"`
	}
}
