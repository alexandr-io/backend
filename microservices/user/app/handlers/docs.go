// Package handlers of User API
//
// Documentation for User API
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
	"github.com/Alexandr-io/Backend/User/data"
	"github.com/alexandr-io/backend_errors"
)

// A set of data consumed to register a new user.
// swagger:parameters registerUser
type userRegisterParametersWrapper struct {
	// The information to register a new user
	// in: body
	Body userRegister
}

// A user return in a response.
// swagger:response userResponse
type userResponseWrapper struct {
	// A single user
	// in: body
	Body data.User
}

// A bad request error response.
// swagger:response badRequestErrorResponse
type badRequestErrorResponseWrapper struct {
	// A list of described bad request error
	// in: body
	Body backend_errors.BadInput
}
