package errorTypes

type BadInput struct {
	JsonError []byte
	Err       error
}

func (e *BadInput) Error() string {
	return e.Err.Error()
}

func (e *BadInput) Unwrap() error { return e.Err }
