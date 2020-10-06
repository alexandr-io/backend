package data

// BadInput is the error type for the BadRequest error.
type BadInput struct {
	JSONError []byte
	Err       error
}

// Error return the string error message contained in the error.
func (e *BadInput) Error() string {
	return e.Err.Error()
}

// Unwrap unwrap the error type BadInput.
func (e *BadInput) Unwrap() error { return e.Err }
