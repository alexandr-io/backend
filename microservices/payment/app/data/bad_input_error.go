package data

// BadInputError is the error type for the BadRequest error.
type BadInputError struct {
	JSONError []byte
	Err       error
}

// Error return the string error message contained in the error.
func (e *BadInputError) Error() string {
	return e.Err.Error()
}

// Unwrap unwraps the error type BadInput.
func (e *BadInputError) Unwrap() error { return e.Err }
