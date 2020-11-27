package data

// BadInput is the struct used to format json in case of bad input
type BadInput struct {
	Fields []Field `json:"fields"`
}

// Field is a nested data contained in BadInput
type Field struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

// ErrorType is a string type to contain error types
type ErrorType string

// error type variables
const (
	Username ErrorType = "username"
	Required ErrorType = "required"
)

// ErrorTypes is a map with an ErrorType as a key and the value that should be used for this error
var ErrorTypes = map[ErrorType]string{
	Username: "The username is invalid",
	Required: "The field is required",
}
