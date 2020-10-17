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
	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/berrors"
)

//
// Input
//

// A set of data consumed to register a new user
// swagger:parameters register
type userRegisterParametersWrapper struct {
	// The information to register a new user
	// in: body
	Body data.UserRegister
}

// A set of data consumed to login a user
// swagger:parameters login
type userLoginParametersWrapper struct {
	// The information to login a user
	// in: body
	Body data.UserLogin
}

// A set of data consumed to refresh an auth and refresh token
// swagger:parameters refresh_token
type authRefreshParametersWrapper struct {
	// The information to refresh an auth and refresh token
	// in: body
	Body authRefresh
}

//
// Output
//

// User data with auth and refresh token
// swagger:response userResponse
type userResponseWrapper struct {
	// A single user
	// in: body
	Body data.User
}

//// A simple response from a simple route
//// swagger:response authResponse
//type authResponseWrapper struct {
//	// A username gotten from the jwt
//	// in: body
//	Body struct {
//		// The connected user's username
//		// Required: true
//		// Example: john
//		Username string `json:"username"`
//	}
//}

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
