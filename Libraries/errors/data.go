package errors

type badInput struct {
	Fields []field `json:"fields"`
}

type field struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}
